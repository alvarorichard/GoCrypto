package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User is the user model
type User struct {
	Id        string    `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Email     string    `gorm:"uniqueIndex" json:"email" binding:"required"`
	Name      string    `json:"name" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Type      string
	Coins     []Coin `gorm:"foreignKey:OwnerID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Id = uuid.New().String()
	u.CreatedAt = time.Now()
	return
}
