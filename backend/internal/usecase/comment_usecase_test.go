package usecase_test

import (
	"testing"

	"example.com/interview-question-009/internal/domain"
	"example.com/interview-question-009/internal/repository"
	"example.com/interview-question-009/internal/testutil"
	"example.com/interview-question-009/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func seedPostComment(t *testing.T, db *gorm.DB) domain.Post {
	t.Helper()
	post := domain.Post{Author: "Change can", ImageURL: "http://example.com/img.jpg"}
	require.NoError(t, db.Create(&post).Error)
	return post
}

func TestListComments(t *testing.T) {
	db := testutil.SetupTestDB(t)
	uc := usecase.NewCommentUseCase(repository.NewCommentRepository(db))
	post := seedPostComment(t, db)

	db.Create(&domain.Comment{PostID: post.ID, Author: "Blend 285", Content: "hello"})
	db.Create(&domain.Comment{PostID: post.ID, Author: "Blend 285", Content: "world"})

	comments, err := uc.ListComments(post.ID)

	require.NoError(t, err)
	assert.Len(t, comments, 2)
}

func TestAddComment(t *testing.T) {
	db := testutil.SetupTestDB(t)
	uc := usecase.NewCommentUseCase(repository.NewCommentRepository(db))
	post := seedPostComment(t, db)

	comment, err := uc.AddComment(post.ID, "have a good day", "Blend 285")

	require.NoError(t, err)
	assert.NotZero(t, comment.ID)
	assert.Equal(t, "Blend 285", comment.Author)
	assert.Equal(t, "have a good day", comment.Content)
	assert.Equal(t, post.ID, comment.PostID)
}

func TestAddComment_EmptyContent(t *testing.T) {
	db := testutil.SetupTestDB(t)
	uc := usecase.NewCommentUseCase(repository.NewCommentRepository(db))
	post := seedPostComment(t, db)

	// empty content passes at use-case level; validation lives in handler
	comment, err := uc.AddComment(post.ID, "", "Blend 285")

	require.NoError(t, err)
	assert.Equal(t, "", comment.Content)
}

func TestRemoveComment(t *testing.T) {
	db := testutil.SetupTestDB(t)
	uc := usecase.NewCommentUseCase(repository.NewCommentRepository(db))
	post := seedPostComment(t, db)

	comment, err := uc.AddComment(post.ID, "to delete", "Blend 285")
	require.NoError(t, err)

	err = uc.RemoveComment(comment.ID)
	require.NoError(t, err)

	remaining, _ := uc.ListComments(post.ID)
	assert.Empty(t, remaining)
}
