package token

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func (t *Token) GenerateToken() error {
	expireAt := time.Now().Add(time.Hour * 24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": t.Username,
			"password": t.Password,
			"exp":      expireAt.Unix(),
		})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_TOKEN_SECRET_KEY")))
	if err != nil {
		return err
	}

	t.Token = tokenString
	t.ExpiresAt = &expireAt
	return nil
}
