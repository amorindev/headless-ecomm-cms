package handler

import (
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/logger"
	"github.com/amorindev/headless-ecomm-cms/cmd/api/middlewares"
	categoryP "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/categories/port"
	productP "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/port"
	varOptP "github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/var-options/port"
)

type Handler struct {
	ProductSrv    productP.ProductSrv
	CategoryRepo  categoryP.CategoryRepo
	VarOptionRepo varOptP.VariationOptionRepo
}

func NewProductHdl(
	server *http.ServeMux,
	productSrv productP.ProductSrv,
	categoryRepo categoryP.CategoryRepo,
	varOptionRepo varOptP.VariationOptionRepo,
	authMdw *middlewares.AuthMiddleware,
) *Handler {
	h := &Handler{
		ProductSrv:    productSrv,
		CategoryRepo:  categoryRepo,
		VarOptionRepo: varOptionRepo,
	}

	server.HandleFunc("GET /products", h.GetAll)
	server.HandleFunc("POST /products/load-from-zip", logger.LoggerMdw(h.LoadFromZip))

	return h
}
