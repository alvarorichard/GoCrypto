package carteira

import (
	"errors"
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v4"
    "gocrypto_back/types"
    "gorm.io/gorm"
)

func getUserIDFromToken(c *fiber.Ctx) (uint, error) {
    user := c.Locals("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    userID, ok := claims["user_id"].(float64)
    if !ok {
        return 0, errors.New("invalid user_id in token")
    }
    return uint(userID), nil
}

func parseOperation(c *fiber.Ctx) (types.Operation, error) {
    var operation types.Operation
    if err := c.BodyParser(&operation); err != nil {
        return operation, err
    }
    // Add validation logic for operation here if needed
    return operation, nil
}

func Buy(c *fiber.Ctx) error {
    userID, err := getUserIDFromToken(c)
    if err != nil {
        return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
    }

    operation, err := parseOperation(c)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"message": "Bad Request"})
    }

    db := c.Locals("db").(*gorm.DB)
    balance := types.Balance{
        UserID: userID,
        CoinID: operation.CoinID,
        Amount: operation.Amount,
        Operation: "buy",
    }

    if err := db.Create(&balance).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"message": "Internal Server Error"})
    }

    return c.JSON(fiber.Map{"message": "Coin purchased successfully"})
}

func Sell(c *fiber.Ctx) error {
    userID, err := getUserIDFromToken(c)
    if err != nil {
        return c.Status(401).JSON(fiber.Map{"message": "Unauthorized"})
    }

    operation, err := parseOperation(c)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"message": "Bad Request"})
    }

    db := c.Locals("db").(*gorm.DB)
    balance := types.Balance{
        UserID: userID,
        CoinID: operation.CoinID,
        Amount: operation.Amount,
        Operation: "sell",
    }

    if err := db.Create(&balance).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"message": "Internal Server Error"})
    }

    return c.JSON(fiber.Map{"message": "Coin sold successfully"})
}
