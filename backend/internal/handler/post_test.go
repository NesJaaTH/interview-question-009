package handler_test

import (
	"encoding/json"
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

func newPostApp(db *gorm.DB) *fiber.App {
	h := handler.NewPostHandler(usecase.NewPostUseCase(repository.NewPostRepository(db)))
	app := fiber.New()
	app.Get("/:postID", h.GetPost)
	return app
}

func TestGetPost_OK(t *testing.T) {
	db := testutil.SetupTestDB(t)
	seed := domain.Post{Author: "Change can", ImageURL: "http://example.com/img.jpg"}
	require.NoError(t, db.Create(&seed).Error)

	resp, err := newPostApp(db).Test(httptest.NewRequest("GET", "/1", nil))

	require.NoError(t, err)
	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var post domain.Post
	require.NoError(t, json.NewDecoder(resp.Body).Decode(&post))
	assert.Equal(t, "Change can", post.Author)
}

func TestGetPost_NotFound(t *testing.T) {
	db := testutil.SetupTestDB(t)

	resp, err := newPostApp(db).Test(httptest.NewRequest("GET", "/999", nil))

	require.NoError(t, err)
	assert.Equal(t, fiber.StatusNotFound, resp.StatusCode)
}

func TestGetPost_InvalidID(t *testing.T) {
	db := testutil.SetupTestDB(t)

	resp, err := newPostApp(db).Test(httptest.NewRequest("GET", "/abc", nil))

	require.NoError(t, err)
	assert.Equal(t, fiber.StatusBadRequest, resp.StatusCode)
}
