package handler

import (
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/logger"
	"github.com/amorindev/headless-ecomm-cms/cmd/api/middlewares"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/addresses/port"
)

type Handler struct {
	AddressSrv port.AddressSrv
}

func NewAddressHdl(server *http.ServeMux, addressSrv port.AddressSrv, authMdw *middlewares.AuthMiddleware) *Handler {
	h := &Handler{
		AddressSrv: addressSrv,
	}

	getAllH := logger.LoggerMdw(authMdw.AccessTokenMdw(h.GetAll))
	createH := logger.LoggerMdw(authMdw.AccessTokenMdw(h.Create))

	server.HandleFunc("GET /addresses", getAllH)
	server.HandleFunc("POST /addresses", createH)
	
	return h
}
