package testutil

import (
	"testing"

	"example.com/interview-question-009/internal/domain"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// SetupTestDB opens an isolated in-memory SQLite DB and migrates all models.
// Returns *gorm.DB so each test constructs its own repositories — no global state.
func SetupTestDB(t *testing.T) *gorm.DB {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("testutil: failed to open in-memory db: %v", err)
	}

	if err := db.AutoMigrate(&domain.Post{}, &domain.Comment{}, &domain.User{}); err != nil {
		t.Fatalf("testutil: failed to migrate: %v", err)
	}

	t.Cleanup(func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	})

	return db
}
