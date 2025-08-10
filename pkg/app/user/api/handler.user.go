package handler

import (
	"net/http"

	"com.fernando/cmd/api/logger"
	mdw "com.fernando/cmd/api/middlewares"
	"com.fernando/pkg/app/users/port"
)

type Handler struct {
	UserSrv port.UserSrv
}

func NewHandler(server *http.ServeMux, userSrv port.UserSrv, authMdw *mdw.AuthMiddleware) *Handler {
	h := &Handler{
		UserSrv: userSrv,
	}

	userH := logger.LoggerMdw(authMdw.AccessTokenMdw(h.UserHandler))
	server.HandleFunc("GET /v1/users/me", userH)

	return h
}