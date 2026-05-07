package repository

import (
	"example.com/interview-question-009/internal/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository returns a domain.UserRepository backed by the given *gorm.DB.
func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

// GetByUsername fetches the user row with the given username.
// Returns gorm.ErrRecordNotFound when no row matches.
func (r *userRepository) GetByUsername(username string) (*domain.User, error) {
	var user domain.User
	result := r.db.Where("username = ?", username).First(&user)
	return &user, result.Error
}
