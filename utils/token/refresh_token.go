package token

func (t *Token) RefreshToken() {
	t.GenerateToken()
}
