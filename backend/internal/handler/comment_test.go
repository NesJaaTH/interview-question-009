package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"example.com/interview-question-009/internal/domain"
	"example.com/interview-question-009/internal/handler"
	"example.com/interview-question-009/internal/repository"
	"example.com/interview-question-009/internal/testutil"
	"example.com/interview-question-009/internal/usecase"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// injectAuth simulates what JWTProtected middleware does in production by
// setting displayName in c.Locals before the handler runs.
func injectAuth(c fiber.Ctx) error {
	c.Locals("displayName", "Blend 285")
	return c.Next()
}

func newCommentApp(db *gorm.DB) *fiber.App {
	h := handler.NewCommentHandler(usecase.NewCommentUseCase(repository.NewCommentRepository(db)))
	app := fiber.New()
	app.Get("/:postID/comments", h.GetComments)
	app.Post("/:postID/comments", injectAuth, h.CreateComment)
	app.Delete("/:postID/comments/:commentID", injectAuth, h.DeleteComment)
	return app
}

func seedPostHandler(t *testing.T, db *gorm.DB) domain.Post {
	t.Helper()
	post := domain.Post{Author: "Change can", ImageURL: "http://example.com/img.jpg"}
	require.NoError(t, db.Create(&post).Error)
	return post
}

func TestGetComments_OK(t *testing.T) {
	db := testutil.SetupTestDB(t)
	post := seedPostHandler(t, db)
	db.Create(&domain.Comment{PostID: post.ID, Author: "Blend 285", Content: "hello"})
	db.Create(&domain.Comment{PostID: post.ID, Author: "Blend 285", Content: "world"})

	resp, err := newCommentApp(db).Test(httptest.NewRequest("GET", fmt.Sprintf("/%d/comments", post.ID), nil))

	require.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var comments []domain.Comment
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&comments))
	assert.Len(t, comments, 2)
	assert.Equal(t, "hello", comments[0].Content)
}

func TestGetComments_Empty(t *testing.T) {
	db := testutil.SetupTestDB(t)
	post := seedPostHandler(t, db)

	resp, err := newCommentApp(db).Test(httptest.NewRequest("GET", fmt.Sprintf("/%d/comments", post.ID), nil))

	require.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var comments []domain.Comment
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&comments))
	assert.Empty(t, comments)
}

func TestCreateComment_OK(t *testing.T) {
	db := testutil.SetupTestDB(t)
	post := seedPostHandler(t, db)

	body, _ := json.Marshal(map[string]string{"content": "have a good day"})
	req := httptest.NewRequest("POST", fmt.Sprintf("/%d/comments", post.ID), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := newCommentApp(db).Test(req)

	require.NoError(t, err)
	assert.Equal(t, fiber.StatusCreated, resp.StatusCode)

	var comment domain.Comment
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&comment))
	assert.NotZero(t, comment.ID)
	assert.Equal(t, "Blend 285", comment.Author)
	assert.Equal(t, "have a good day", comment.Content)
}

func TestCreateComment_EmptyContent(t *testing.T) {
	db := testutil.SetupTestDB(t)
	post := seedPostHandler(t, db)

	body, _ := json.Marshal(map[string]string{"content": ""})
	req := httptest.NewRequest("POST", fmt.Sprintf("/%d/comments", post.ID), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := newCommentApp(db).Test(req)

	require.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestCreateComment_InvalidBody(t *testing.T) {
	db := testutil.SetupTestDB(t)
	post := seedPostHandler(t, db)

	req := httptest.NewRequest("POST", fmt.Sprintf("/%d/comments", post.ID), bytes.NewReader([]byte("not-json")))
	req.Header.Set("Content-Type", "application/json")
	resp, err := newCommentApp(db).Test(req)

	require.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}

func TestDeleteComment_OK(t *testing.T) {
	db := testutil.SetupTestDB(t)
	post := seedPostHandler(t, db)
	comment := domain.Comment{PostID: post.ID, Author: "Blend 285", Content: "to delete"}
	require.NoError(t, db.Create(&comment).Error)

	resp, err := newCommentApp(db).Test(httptest.NewRequest("DELETE", fmt.Sprintf("/%d/comments/%d", post.ID, comment.ID), nil))

	require.NoError(t, err)
	assert.Equal(t, fiber.StatusNoContent, resp.StatusCode)
}

func TestDeleteComment_InvalidID(t *testing.T) {
	db := testutil.SetupTestDB(t)
	post := seedPostHandler(t, db)

	resp, err := newCommentApp(db).Test(httptest.NewRequest("DELETE", fmt.Sprintf("/%d/comments/abc", post.ID), nil))

	require.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}
