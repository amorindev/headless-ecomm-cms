package auth

type TokenSrv struct {
	AccessSecret  string
	RefreshSecret string
	Issuer        string
}

func NewTokenSrv(accessSecret, refreshSecret, issuer string) *TokenSrv {
	return &TokenSrv{
		AccessSecret:  accessSecret,
		RefreshSecret: refreshSecret,
		Issuer:        issuer,
	}
}

func (ts *TokenSrv) NewAccessToken(userID string, email string, roles []string) (*AccessTokenClaims, string, error) {
	claims := NewAccessTokenClaim(userID, email, ts.Issuer, roles)
	token, err := claims.GetToken(ts.AccessSecret)
	if err != nil {
		return nil, "", err
	}
	return claims, token, nil
}

func (ts *TokenSrv) ParseAccessToken(tokenString string) (*AccessTokenClaims, error) {
	return GetAccessTokenFromJWT(tokenString, ts.AccessSecret)
}

func (ts *TokenSrv) NewRefreshToken(userID string, rememberMe bool) (*RefreshTokenClaims, string, error) {
	claims := NewRefreshTokenClaim(userID, rememberMe)
	token, err := claims.GetToken(ts.RefreshSecret)
	if err != nil {
		return nil, "", err
	}
	return claims, token, nil
}

func (ts *TokenSrv) ParseRefreshToken(tokenString string) (*RefreshTokenClaims, error) {
	return GetRefreshTokenFromJWT(tokenString, ts.RefreshSecret)
}
