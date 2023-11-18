package middlewares

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var JwtSecret = []byte("GOCRYPTOROCKS")

func GenerateJWT(userID uint, isAdmin bool) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["iat"] = time.Now().Unix()
	claims["admin"] = isAdmin
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token
	t, err := token.SignedString(JwtSecret)
	return t, err
}
