// Package repository contains GORM-backed implementations of the domain repository interfaces.
package repository

import (
	"example.com/interview-question-009/internal/domain"
	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

// NewPostRepository returns a domain.PostRepository backed by the given *gorm.DB.
func NewPostRepository(db *gorm.DB) domain.PostRepository {
	return &postRepository{db: db}
}

// GetByID fetches the post with the given primary key.
// Returns gorm.ErrRecordNotFound when no row matches.
func (r *postRepository) GetByID(id uint) (*domain.Post, error) {
	var post domain.Post
	result := r.db.First(&post, id)
	return &post, result.Error
}
