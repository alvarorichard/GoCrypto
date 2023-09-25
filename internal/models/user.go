package models

import "gorm.io/gorm"

// User is the user model
type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex" json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Type     string
	Coins    []Coin `gorm:"many2many:user_coins;"`
}

// Save saves the user to the database
func (u *User) Save(db *gorm.DB) (err error) {
	err = db.Create(u).Error
	return
}
