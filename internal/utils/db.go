package utils

import (
	"github.com/alvarorichard/GoCrypto/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewDB creates a new database connection
func NewDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
}

// MigrateDB migrates the database
func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{}, &models.Coin{})
}
