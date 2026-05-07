package repository_test

import (
	"testing"

	"example.com/interview-question-009/internal/domain"
	"example.com/interview-question-009/internal/repository"
	"example.com/interview-question-009/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func seedPostRepo(t *testing.T, db *gorm.DB) domain.Post {
	t.Helper()
	post := domain.Post{Author: "Change can", ImageURL: "http://example.com/img.jpg"}
	require.NoError(t, db.Create(&post).Error)
	return post
}

func TestGetCommentsByPostID(t *testing.T) {
	db := testutil.SetupTestDB(t)
	repo := repository.NewCommentRepository(db)
	post := seedPostRepo(t, db)

	other := domain.Post{Author: "Other", ImageURL: "http://example.com/other.jpg"}
	require.NoError(t, db.Create(&other).Error)

	db.Create(&domain.Comment{PostID: post.ID, Author: "Blend 285", Content: "first"})
	db.Create(&domain.Comment{PostID: post.ID, Author: "Blend 285", Content: "second"})
	db.Create(&domain.Comment{PostID: other.ID, Author: "Blend 285", Content: "other post"})

	comments, err := repo.GetByPostID(post.ID)

	require.NoError(t, err)
	assert.Len(t, comments, 2)
	assert.Equal(t, "first", comments[0].Content)
	assert.Equal(t, "second", comments[1].Content)
}

func TestGetCommentsByPostID_Empty(t *testing.T) {
	db := testutil.SetupTestDB(t)
	repo := repository.NewCommentRepository(db)
	post := seedPostRepo(t, db)

	comments, err := repo.GetByPostID(post.ID)

	require.NoError(t, err)
	assert.Empty(t, comments)
}

func TestCreateComment(t *testing.T) {
	db := testutil.SetupTestDB(t)
	repo := repository.NewCommentRepository(db)
	post := seedPostRepo(t, db)

	comment := &domain.Comment{PostID: post.ID, Author: "Blend 285", Content: "hello"}
	err := repo.Create(comment)

	require.NoError(t, err)
	assert.NotZero(t, comment.ID)
	assert.False(t, comment.CreatedAt.IsZero())
}

func TestDeleteCommentByID(t *testing.T) {
	db := testutil.SetupTestDB(t)
	repo := repository.NewCommentRepository(db)
	post := seedPostRepo(t, db)

	comment := &domain.Comment{PostID: post.ID, Author: "Blend 285", Content: "to delete"}
	require.NoError(t, db.Create(comment).Error)

	err := repo.DeleteByID(comment.ID)
	require.NoError(t, err)

	remaining, _ := repo.GetByPostID(post.ID)
	assert.Empty(t, remaining)
}
