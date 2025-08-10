package middlewares

import "com.fernando/internal/auth"

type AuthMiddleware struct {
	AuthSrv *auth.TokenSrv
}

func NewAuthMdw(authSrv *auth.TokenSrv) *AuthMiddleware {
	return &AuthMiddleware{
		AuthSrv: authSrv,
	}
}
