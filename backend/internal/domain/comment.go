package domain

import "time"

// Comment represents a user comment attached to a Post.
type Comment struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PostID    uint      `gorm:"not null;index"           json:"post_id"`
	Author    string    `gorm:"not null"                 json:"author"`
	Content   string    `gorm:"not null"                 json:"content"`
	CreatedAt time.Time `                                json:"created_at"`
}

// CommentRepository is the persistence contract for Comment.
// The concrete implementation lives in the repository package.
type CommentRepository interface {
	// GetByPostID returns all comments for the given post, ordered by created_at ASC.
	GetByPostID(postID uint) ([]Comment, error)
	// Create persists a new comment and populates its auto-generated fields (ID, CreatedAt).
	Create(comment *Comment) error
	// DeleteByID hard-deletes the comment with the given primary key.
	DeleteByID(id uint) error
}
