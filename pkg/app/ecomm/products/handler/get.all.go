package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/message"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/products/core"
)

func (h Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()

	pageStr := query.Get("page")
	limitStr := query.Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1|| limit > 100 {
		limit = 10
	}

	products, count, totalPages, err := h.ProductSrv.GetAll(r.Context(), int64(limit), int64(page))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	// localhost:8000
	host := r.Host
	// /products
	basePath := "/v1" + r.URL.Path

	var nextURL *string
	if int64(page) < totalPages {
		url := fmt.Sprintf("%s://%s%s?page=%d&limit=%d", scheme, host, basePath, page+1, limit)
		nextURL = &url
	}

	var prevURL *string
	if page > 1 {
		url := fmt.Sprintf("%s://%s%s?page=%d&limit=%d", scheme, host, basePath, page-1, limit)
		prevURL = &url
	}

	resp := &core.ProductResp{
		Count:    count,
		Pages:    totalPages,
		Next:     nextURL,
		Previous: prevURL,
		Products: products,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

