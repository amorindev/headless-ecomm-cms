package handler

import (
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/logger"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/admin/api"
)

type Handler struct {
	Renderer   *api.TemplateRenderer
	ApiBaseUrl string
}

func NewAdminHandler(mux *http.ServeMux, tr *api.TemplateRenderer, apiBaseUrl string) *Handler {
	h := &Handler{
		Renderer:   tr,
		ApiBaseUrl: apiBaseUrl,
	}

	// TODO ver que nombres poner
	// * Templates

	mux.HandleFunc("/admin/sign-in-admin", logger.LoggerMdw(h.SignInAdminPage))
	mux.HandleFunc("/admin/sign-up-admin", h.SignUpAdminPage)
	mux.HandleFunc("/admin/products", h.ProductsPage)
	mux.HandleFunc("/admin/categories", h.CategoriesPage)
	mux.HandleFunc("/admin/onboardings", h.OnboardingsPage)
	mux.HandleFunc("/admin/roles", h.RolesPage)

	// * Api
	mux.HandleFunc("POST /sign-up-admin/submit", h.SignUpAdmin)
	mux.HandleFunc("POST /sign-in-admin/submit", h.SignInAdmin)

	return h
}
