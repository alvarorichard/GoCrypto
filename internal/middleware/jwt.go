package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	UserID   int
	UserName string
	Admin    bool
}

func New(id int, userName string) *Jwt {
	return &Jwt{
		UserID:   id,
		UserName: userName,
	}
}

func (j *Jwt) Sign() (token string, err error) {
	t := jwt.New(jwt.SigningMethodHS256)
	claims := t.Claims.(jwt.MapClaims)
	claims["id"] = j.UserID
	claims["name"] = j.UserName
	claims["admin"] = j.Admin
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	return t.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func VerifyToken(token string) (claims jwt.MapClaims, err error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return t.Claims.(jwt.MapClaims), nil
}
