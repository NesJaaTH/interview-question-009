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

func seedPostUC(t *testing.T, db *gorm.DB) domain.Post {
	t.Helper()
	post := domain.Post{Author: "Change can", ImageURL: "http://example.com/img.jpg"}
	require.NoError(t, db.Create(&post).Error)
	return post
}

func TestGetPost(t *testing.T) {
	db := testutil.SetupTestDB(t)
	uc := usecase.NewPostUseCase(repository.NewPostRepository(db))
	post := seedPostUC(t, db)

	result, err := uc.GetPost(post.ID)

	require.NoError(t, err)
	assert.Equal(t, post.ID, result.ID)
	assert.Equal(t, "Change can", result.Author)
}

func TestGetPost_NotFound(t *testing.T) {
	db := testutil.SetupTestDB(t)
	uc := usecase.NewPostUseCase(repository.NewPostRepository(db))

	_, err := uc.GetPost(999)

	assert.Error(t, err)
}
