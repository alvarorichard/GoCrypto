package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"gocrypto_back/middlewares"
	"gocrypto_back/routes/admin"
	"gocrypto_back/routes/carteira"
	"gocrypto_back/routes/myUser"
	"gocrypto_back/routes/user"
)

func ApiRouter(api fiber.Router) {
	userGroup := api.Group("/user")
	UserRouter(userGroup)
	api.Use(jwtware.New(jwtware.Config{
		SigningKey: middlewares.JwtSecret,
	}))
	myUserGroup := userGroup.Group("/me")
	MyUserRouter(myUserGroup)
	walletGroup := api.Group("/wallet")
	WalletRouter(walletGroup)
	adminGroup := api.Group("/admin")
	AdminRouter(adminGroup)

}

func UserRouter(group fiber.Router) {
	group.Post("/register", user.Register)
	group.Post("/login", user.Login)
}

func WalletRouter(group fiber.Router) {
	group.Get("/my", carteira.Me)
	group.Post("/buy", carteira.Buy)
	group.Post("/sell", carteira.Sell)
}

func MyUserRouter(group fiber.Router) {
	group.Get("/", myUser.Me)
}

func AdminRouter(group fiber.Router) {
	group.Post("/coin", admin.CreateCoin)
	group.Get("/coin", admin.ListCoins)
}
