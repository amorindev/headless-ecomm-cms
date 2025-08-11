package handler

import (
	"net/http"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/variations/port"

)

type Handler struct {
	VariationSrv port.VariationSrv
}

func NewHandler(server *http.ServeMux, variationSrv port.VariationSrv) *Handler {
	h := &Handler{
		VariationSrv: variationSrv,
	}

	// * Variations
	server.HandleFunc("GET /variations", h.GetAllWithVarOptions)
	server.HandleFunc("POST /variations", h.Create)
	server.HandleFunc("DELETE /variations/{id}", h.Delete)

	// * Variation options
	server.HandleFunc("POST /var-options/{id}/options", h.CreateVarOption)
	server.HandleFunc("DELETE /var-options/{id}", h.DeleteVarOption)

	return h
}
