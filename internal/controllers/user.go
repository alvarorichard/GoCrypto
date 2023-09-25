package controllers

import (
	"fmt"

	"github.com/alvarorichard/GoCrypto/internal/models"
	"gorm.io/gorm"
)

// NewUser creates a new user
func NewUser() (u models.User) {
	u = models.User{
		Type: "user",
	}
	return
}

// Save saves the user to the database
func ListAllUsers(db *gorm.DB) (users []models.User, err error) {
	err = db.Find(&users).Error
	return
}

func DeleteUser(db *gorm.DB, id string) (err error) {
	fmt.Println(id)
	return db.Delete(&models.User{}, id).Error
}
func GetUserByEmail(db *gorm.DB, email string) (user models.User, err error) {
	err = db.Where("email = ?", email).First(&user).Error
	return
}
