package carteira

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"gocrypto_back/types"
	"gorm.io/gorm"
)

func Buy(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)
	db := c.Locals("db").(*gorm.DB)
	operation := types.Operation{}
	err := c.BodyParser(&operation)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad Request"})
	}

	b := types.Balance{
		UserID: uint(userID),
		CoinID: operation.CoinID,
		Amount: operation.Amount,
	}
	err = db.Create(&b).Error
   
	

	if err != nil {

		return c.Status(500).JSON(fiber.Map{"message": "Internal Server Error"})

	}

	return c.JSON(fiber.Map{
		"message": "Coin created",
	})
}

func Sell(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(float64)
	db := c.Locals("db").(*gorm.DB)
	operation := types.Operation{}
	err := c.BodyParser(&operation)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "Bad Request"})
	}

	b := types.Balance{
		UserID: uint(userID),
		CoinID: operation.CoinID,
		Amount: operation.Amount,
	}

	if db.Create(&b).Error == nil {
		return c.Status(500).JSON(fiber.Map{"message": "Internal Server Error"})
	}

	return c.JSON(fiber.Map{
		"message": "Coin created",
	})
}
