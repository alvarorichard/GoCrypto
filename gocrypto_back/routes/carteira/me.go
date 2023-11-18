package carteira

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gocrypto_back/types"
	"gorm.io/gorm"
)

func Me(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)
	db := c.Locals("db").(*gorm.DB)
	u := types.User{}
	db.Where("id = ?", userID).First(&u)
	b := u.Balance
	return c.JSON(b)
}
