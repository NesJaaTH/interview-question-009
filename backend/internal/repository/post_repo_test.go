package repository_test

import (
	"testing"

	"example.com/interview-question-009/internal/domain"
	"example.com/interview-question-009/internal/repository"
	"example.com/interview-question-009/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPostByID(t *testing.T) {
	db := testutil.SetupTestDB(t)
	repo := repository.NewPostRepository(db)

	seed := domain.Post{Author: "Change can", ImageURL: "http://example.com/img.jpg"}
	require.NoError(t, db.Create(&seed).Error)

	post, err := repo.GetByID(seed.ID)

	require.NoError(t, err)
	assert.Equal(t, seed.ID, post.ID)
	assert.Equal(t, "Change can", post.Author)
}

func TestGetPostByID_NotFound(t *testing.T) {
	db := testutil.SetupTestDB(t)
	repo := repository.NewPostRepository(db)

	_, err := repo.GetByID(999)

	assert.Error(t, err)
}
