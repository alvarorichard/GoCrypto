package main

import (
	"github.com/gofiber/fiber/v2"
	"gocrypto_back/routes"
	"gocrypto_back/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&types.User{}, &types.Coin{}, &types.Balance{})
	if err != nil {
		panic("failed to migrate database")
	}

	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	api := app.Group("/api")
	routes.ApiRouter(api)
	_ = app.Listen(":3000")
}
