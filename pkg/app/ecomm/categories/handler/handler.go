package handler

import (
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/categories/port"
)

type Handler struct {
	CategorySrv port.CategorySrv
}

func NewCategoryHdl(server *http.ServeMux, categorySrv port.CategorySrv) *Handler {
	h := &Handler{
		CategorySrv: categorySrv,
	}

	server.HandleFunc("GET /categories", h.GetAll)
	server.HandleFunc("POST /categories", h.Create)
	server.HandleFunc("PUT /categories/{id}", h.Update)
	server.HandleFunc("DELETE /categories/{id}", h.Delete)

	return h
}
