package views

import (
	"fmt"

	"github.com/alvarorichard/GoCrypto/internal/controllers"
	"github.com/alvarorichard/GoCrypto/internal/middleware"
	"github.com/alvarorichard/GoCrypto/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type CoinView struct {
	db *gorm.DB
}

// NewCoinView creates a new coin view
func NewCoinView(db *gorm.DB) (c CoinView) {
	c = CoinView{
		db: db,
	}
	return c
}

func (c *CoinView) Register(r *gin.Engine) {
	r.POST("/coins/new", middleware.Auth(), c.Create)
	r.GET("/coins", middleware.Auth(), c.ListCoinsFromUser)
}

func (cv *CoinView) Create(c *gin.Context) {
	claims, _ := c.Get("claims")
	fmt.Println(claims)
	j := claims.(jwt.MapClaims)
	idUint := j["id"].(string)
	println(idUint)
	var coin models.Coin
	if err := c.ShouldBindJSON(&coin); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	coin.OwnerID = idUint
	err := cv.db.Create(&coin).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	}

	c.JSON(200, coin)
}

func (cv *CoinView) ListCoinsFromUser(c *gin.Context) {
	claims, _ := c.Get("claims")
	fmt.Println(claims)
	j := claims.(jwt.MapClaims)
	idUint := j["id"].(string)
	coins, err := controllers.ListCoinsFromUser(cv.db, idUint)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
	}
	c.JSON(200, gin.H{
		"coins": coins,
	})
}
