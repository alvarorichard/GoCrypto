package carteira

import (
	"errors"
    "github.com/gofiber/fiber/v2"
    "github.com/golang-jwt/jwt/v4"
    "gocrypto_back/types"
    "gorm.io/gorm"
)

func Me(c *fiber.Ctx) error {
    user := c.Locals("user").(*jwt.Token)
    claims := user.Claims.(jwt.MapClaims)
    userID := uint(claims["user_id"].(float64)) // Convert to uint

    db := c.Locals("db").(*gorm.DB)
    var u types.User

    result := db.Preload("Balance").Where("id = ?", userID).First(&u)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return c.Status(404).JSON(fiber.Map{"message": "User not found"})
        }
        return c.Status(500).JSON(fiber.Map{"message": "Internal Server Error"})
    }

    if u.Balance == nil {
        // Handle case where balance is null
        return c.JSON(fiber.Map{"balance": 0}) // or any appropriate default value or message
    }

    var res []types.WalletResponse

    for _, b := range u.Balance {
        var c types.Coin
        db.Where("id = ?", b.CoinID).First(&c)
        r := types.WalletResponse{
            Date: b.CreatedAt,
            Ammount: b.Amount,
            CoinName: c.Name,
            Operation: b.Operation,
        }
        res = append(res, r)
    }




    return c.JSON(res)
}
