package auth

func GenerateToken(id string) (*Token, error) {
	tokenService := NewTokenService()
	tokenService.SetClaims(&AuthClaims{UserId: id})
	tokenString, err := tokenService.CreateToken()
	if err != nil {
		return nil, err
	}
	return &Token{Value: tokenString}, nil
}
