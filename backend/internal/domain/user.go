package domain

import "time"

// User represents an authenticated user account.
type User struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username    string    `gorm:"uniqueIndex;not null"     json:"username"`
	DisplayName string    `gorm:"not null"                 json:"display_name"`
	Password    string    `gorm:"not null"                 json:"-"` // bcrypt hash, never serialised
	CreatedAt   time.Time `                                json:"created_at"`
}

// UserRepository is the persistence contract for User.
type UserRepository interface {
	// GetByUsername returns the user with the given username, or an error if not found.
	GetByUsername(username string) (*User, error)
}
