package types

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string    `json:"email" gorm:"uniqueIndex;not null;type:varchar(100)"`
	Password string    `json:"password"`
	Admin    bool      `json:"admin"`
	Balance  []Balance `gorm:"foreignKey:UserID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.Admin = false
	return
}

type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}


type Promote struct {
	ID uint `json:"user_id"`
}