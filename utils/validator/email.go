package validator

import "net/mail"

type EmailValidator string

func NewEmailValidator(email string) EmailValidator {
	return EmailValidator(email)
}

func (e EmailValidator) Validate() error {
	_, validationErr := mail.ParseAddress(string(e))
	if validationErr != nil {
		return validationErr
	}
	return nil
}