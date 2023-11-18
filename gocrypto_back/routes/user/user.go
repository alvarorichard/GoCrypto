package user

import (
	"github.com/gofiber/fiber/v2"
	"gocrypto_back/middlewares"
	"gocrypto_back/types"
	"gorm.io/gorm"
)

func Register(c *fiber.Ctx) error {
	// get db from fiber context
	db := c.Locals("db").(*gorm.DB)
	message := types.Message{}
	data := new(types.NewUser)
	if err := c.BodyParser(data); err != nil {
		// Handle error
		message.Message = err.Error()
		message.Type = "error"
		m, _ := message.MarshalBinary()
		return c.Status(fiber.StatusBadRequest).Send(m)
	}
	user := types.User{
		Email:    data.Email,
		Password: data.Password,
	}
	err := db.Create(&user).Error
	if err != nil {
		message.Message = err.Error()
		message.Type = "error"
		m, _ := message.MarshalBinary()
		return c.Status(fiber.StatusBadRequest).Send(m)
	}
	jwt, err := middlewares.GenerateJWT(user.ID, user.Admin)
	message.Message = jwt
	message.Type = "success"

	m, _ := message.MarshalBinary()
	return c.Send(m)
}

func Login(c *fiber.Ctx) error {
	db := c.Locals("db").(*gorm.DB)
	message := types.Message{}
	data := new(types.NewUser)
	if err := c.BodyParser(data); err != nil {
		// Handle error
		message.Message = err.Error()
		message.Type = "error"
		m, _ := message.MarshalBinary()
		return c.Status(fiber.StatusBadRequest).Send(m)
	}
	user := types.User{}
	err := db.Where("email = ?", data.Email).First(&user).Error
	if err != nil {
		message.Message = err.Error()
		message.Type = "error"
		m, _ := message.MarshalBinary()
		return c.Status(fiber.StatusBadRequest).Send(m)
	}
	if user.Password != data.Password {
		message.Message = "Incorrect password"
		message.Type = "error"
		m, _ := message.MarshalBinary()
		return c.Status(fiber.StatusBadRequest).Send(m)
	}
	jwt, err := middlewares.GenerateJWT(user.ID, user.Admin)
	message.Message = jwt
	message.Type = "success"
	m, _ := message.MarshalBinary()
	return c.Send(m)
}
