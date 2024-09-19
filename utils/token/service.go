package token

import "time"

type Token struct {
	Token     string     `json:"token"`
	ExpiresAt *time.Time `json:"expires_at"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
}

type TokenService interface {
	NewToken(username string, password string) *Token
	GenerateToken() error
	ValidateToken() error
	RefreshToken()
}
