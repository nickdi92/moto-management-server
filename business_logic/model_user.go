package business_logic

import (
	"time"
)

type User struct {
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Lastname  string     `json:"lastname"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Token     string     `json:"token"`
	ExpireAt  *time.Time `json:"expire_at"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}