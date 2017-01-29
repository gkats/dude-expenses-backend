package auth

func GenerateToken(id string, secret string) (*Token, error) {
	tokenService := newTokenService(secret)
	tokenService.SetClaims(&AuthClaims{UserId: id})
	tokenString, err := tokenService.CreateToken()
	if err != nil {
		return nil, err
	}
	return &Token{Value: tokenString}, nil
}
