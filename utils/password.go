package utils

import "golang.org/x/crypto/bcrypt"

type Password string

func (p Password) Hash() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(p), 14)
	return string(bytes), err
}
