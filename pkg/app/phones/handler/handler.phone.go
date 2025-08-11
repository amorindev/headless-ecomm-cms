package handler

import (
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/logger"
	mdw "github.com/amorindev/headless-ecomm-cms/cmd/api/middlewares"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/phones/port"
)

type Handler struct {
	PhoneSrv port.PhoneSrv
}

func NewPhoneHdl(server *http.ServeMux, phoneSrv port.PhoneSrv, authMdw *mdw.AuthMiddleware) *Handler {
	h := &Handler{
		PhoneSrv: phoneSrv,
	}

	getAllH := logger.LoggerMdw(authMdw.AccessTokenMdw(h.GetAllByUserID))
	createH := logger.LoggerMdw(authMdw.AccessTokenMdw(h.Create))

	server.HandleFunc("GET /phones", getAllH)
	server.HandleFunc("POST /phones", createH)
	return h
}
