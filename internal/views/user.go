package views

import (
	"fmt"

	"github.com/alvarorichard/GoCrypto/internal/controllers"
	"github.com/alvarorichard/GoCrypto/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserView is the user view
type UserView struct {
	db *gorm.DB
}

// NewUserView creates a new user view
func NewUserView(db *gorm.DB) (u UserView) {
	u = UserView{
		db: db,
	}
	return u
}

// @BasePath /api/v1
// @title GoCrypto API
// @description This is the GoCrypto API documentation
// @version 1
// @host localhost:8080
// @BasePath /api/v1
// @schemes http

func (u *UserView) Register(r *gin.Engine) {
	r.POST("/register", u.Create)
	r.GET("/users", u.ListAll)
	r.POST("/login", u.Login)
	r.GET("/users/me", middleware.Auth(), u.Me)
	r.DELETE("/users/:id", u.DeleteUser)
}

// Create creates a new user
func (u *UserView) Create(c *gin.Context) {
	user := controllers.NewUser()
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err = user.Save(u.db)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	token := middleware.New(int(user.ID), user.Name)
	jwt, err := token.Sign()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": jwt})
}

// ListAll lists all users
func (u *UserView) ListAll(c *gin.Context) {
	users, err := controllers.ListAllUsers(u.db)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"users": users})
}

// DeleteUser deletes a user
func (u *UserView) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := controllers.DeleteUser(u.db, id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": fmt.Sprintf("User %s deleted", id)})
}

// Me returns the current user
func (u *UserView) Me(c *gin.Context) {
	claims, _ := c.Get("claims")
	c.JSON(200, gin.H{"claims": claims})
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *UserView) Login(c *gin.Context) {
	// get email and password from body
	body := LoginRequest{}
	err := c.BindJSON(&body)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := controllers.GetUserByEmail(u.db, body.Email)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if user.Password != body.Password {
		c.JSON(400, gin.H{"error": "Invalid password"})
		return
	}
	token := middleware.New(int(user.ID), user.Name)
	jwt, err := token.Sign()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"token": jwt})

}
