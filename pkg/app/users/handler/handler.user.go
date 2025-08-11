package handler

import (
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/logger"
	"github.com/amorindev/headless-ecomm-cms/cmd/api/middlewares"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/users/port"
)

type Handler struct {
	UserSrv port.UserSrv
}

func NewHandler(server *http.ServeMux, userSrv port.UserSrv, authMdw *middlewares.AuthMiddleware) *Handler {
	h := &Handler{
		UserSrv: userSrv,
	}

	userH := logger.LoggerMdw(authMdw.AccessTokenMdw(h.UserHandler))
	server.HandleFunc("GET /v1/users/me", userH)

	return h
}