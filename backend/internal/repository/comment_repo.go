package repository

import (
	"example.com/interview-question-009/internal/domain"
	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

// NewCommentRepository returns a domain.CommentRepository backed by the given *gorm.DB.
func NewCommentRepository(db *gorm.DB) domain.CommentRepository {
	return &commentRepository{db: db}
}

// GetByPostID returns all comments for the given post, ordered oldest-first.
func (r *commentRepository) GetByPostID(postID uint) ([]domain.Comment, error) {
	var comments []domain.Comment
	result := r.db.
		Where("post_id = ?", postID).
		Order("created_at ASC").
		Find(&comments)
	return comments, result.Error
}

// Create inserts a new comment row and populates the auto-generated ID and CreatedAt fields.
func (r *commentRepository) Create(comment *domain.Comment) error {
	return r.db.Create(comment).Error
}

// DeleteByID hard-deletes the comment row with the given primary key.
func (r *commentRepository) DeleteByID(id uint) error {
	return r.db.Delete(&domain.Comment{}, id).Error
}
