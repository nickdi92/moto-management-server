package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func (t *Token) ValidateToken() error {
	token, err := jwt.Parse(t.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_TOKEN_SECRET_KEY")), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
