package token

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (t *Token) ValidateToken(oldUserToken string) error {
	token, err := jwt.Parse(t.Token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_TOKEN_SECRET_KEY")), nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	exp, _ := token.Claims.GetExpirationTime()
	now := time.Now()
	if now.Equal(exp.Time) || now.After(exp.Time) {
		return errors.New("token is expired")
	}

	if oldUserToken != "" && (oldUserToken != t.Token) {
		return errors.New("token mismatch error")
	}

	return nil
}
