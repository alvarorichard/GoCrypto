package controllers

import (
	"github.com/alvarorichard/GoCrypto/internal/models"
	"gorm.io/gorm"
)

func ListCoinsFromUser(db *gorm.DB, userID string) ([]models.Coin, error) {
	var coins []models.Coin
	err := db.Where("owner_id = ?", userID).Find(&coins).Error
	return coins, err
}
