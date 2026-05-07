// Package domain defines the core entities and repository interfaces.
// It is the innermost layer of the Clean Architecture; nothing here imports
// from other internal packages.
package domain

import "time"

// Post represents a social-media post authored by a user.
type Post struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Author    string    `gorm:"not null"                 json:"author"`
	ImageURL  string    `gorm:"not null"                 json:"image_url"`
	CreatedAt time.Time `                                json:"created_at"`
}

// PostRepository is the persistence contract for Post.
// The concrete implementation lives in the repository package.
type PostRepository interface {
	// GetByID returns the Post with the given primary key, or an error if not found.
	GetByID(id uint) (*Post, error)
}
