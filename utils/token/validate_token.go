package token

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

	now := time.Now()
	if now.Equal(*t.ExpiresAt) {
		return errors.New("token is expired")
	}

	return nil
}
