// Package usecase contains the application business-logic layer.
// Use cases depend on domain interfaces and are themselves exposed as interfaces
// so that handlers are never coupled to concrete implementations.
package usecase

import "example.com/interview-question-009/internal/domain"

// PostUseCase defines the application operations available on a Post.
type PostUseCase interface {
	// GetPost retrieves a single post by its primary key.
	GetPost(id uint) (*domain.Post, error)
}

type postUseCase struct {
	repo domain.PostRepository
}

// NewPostUseCase constructs a PostUseCase backed by the given repository.
func NewPostUseCase(repo domain.PostRepository) PostUseCase {
	return &postUseCase{repo: repo}
}

func (u *postUseCase) GetPost(id uint) (*domain.Post, error) {
	return u.repo.GetByID(id)
}
