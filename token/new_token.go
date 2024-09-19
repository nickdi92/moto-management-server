package token

func (t *Token) NewToken(username string, password string) *Token {
	return &Token{
		Username: username,
		Password: password,
	}
}
