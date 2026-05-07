// Package database handles SQLite connection setup, schema migration, and seeding.
package database

import (
	"log"
	"time"

	"example.com/interview-question-009/internal/domain"
	"github.com/glebarez/sqlite"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Init opens the SQLite file, runs AutoMigrate, seeds initial data, and returns
// the *gorm.DB handle. It calls log.Fatalf on any unrecoverable error.
func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	if err := db.AutoMigrate(&domain.Post{}, &domain.Comment{}, &domain.User{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	seedPosts(db)
	seedUsers(db)
	return db
}

// seedPosts inserts the initial post and comment rows.
// It is idempotent: if at least one post already exists the function returns early.
func seedPosts(db *gorm.DB) {
	var count int64
	db.Model(&domain.Post{}).Count(&count)
	if count > 0 {
		return
	}

	post := domain.Post{
		Author:    "Change can",
		ImageURL:  "https://images.unsplash.com/photo-1450778869180-41d0601e046e?w=700",
		CreatedAt: time.Date(2021, 10, 16, 16, 0, 0, 0, time.UTC),
	}
	db.Create(&post)

	db.Create(&domain.Comment{
		PostID:    post.ID,
		Author:    "Blend 285",
		Content:   "have a good day",
		CreatedAt: time.Date(2021, 10, 16, 16, 5, 0, 0, time.UTC),
	})
}

// seedUsers inserts the initial user account.
// It is idempotent: if at least one user already exists the function returns early.
func seedUsers(db *gorm.DB) {
	var count int64
	db.Model(&domain.User{}).Count(&count)
	if count > 0 {
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte("blend285"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("failed to hash seed password: %v", err)
	}
	db.Create(&domain.User{
		Username:    "blend285",
		DisplayName: "Blend 285",
		Password:    string(hash),
	})
}
