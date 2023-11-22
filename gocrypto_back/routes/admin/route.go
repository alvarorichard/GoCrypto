package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gocrypto_back/types"
	"gorm.io/gorm"
)

func CreateCoin(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["admin"].(bool)
	if !userID {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}
	db := c.Locals("db").(*gorm.DB)
	coin := types.Coin{}
	err := c.BodyParser(&coin)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad Request"})
	}
	err = db.Create(&coin).Error
	return c.JSON(fiber.Map{
		"message": "Coin created",
		"coin":    coin,
	})
}

func ListCoins(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["admin"].(bool)
	if !userID {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}
	db := c.Locals("db").(*gorm.DB)
	coins := []types.Coin{}
	err := db.Find(&coins).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Internal Server Error"})
	}
	return c.JSON(coins)
}


func Promote(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["admin"].(bool)
	if !userID {
		return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
	}

	db := c.Locals("db").(*gorm.DB)

	prom := types.Promote{}

	err := c.BodyParser(&prom)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Internal Server Error"})
	}
	err = db.Table("users").Where("id = ?", prom.ID).Update("admin", true).Error

	// err = db.Table("users").Where("id = ?", prom.ID).Update("admin", true).Error
	
	res := types.Message{
		Type: "success",
		Message: "User promoted with ok",
	}
	return c.JSON(res)
}