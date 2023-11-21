package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors" // Import the CORS middleware
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

	// Configure CORS middleware for all routes
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow any origin
		AllowHeaders: "Origin, Content-Type, Accept", // You can specify the headers you want to allow
		// ... other configuration options
	}))

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	api := app.Group("/api")
	routes.ApiRouter(api)
	_ = app.Listen(":3000")
}
