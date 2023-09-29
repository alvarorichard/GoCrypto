package views

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserView is the user view
func RegisterAll(r *gin.Engine, db *gorm.DB) {
	uv := NewUserView(db)
	uv.Register(r)
	cv := NewCoinView(db)
	cv.Register(r)
}
