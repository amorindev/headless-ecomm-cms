package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/message"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/var-options/domain"
)

func (h Handler) CreateVarOption(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 4 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	type reqBody struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}

	var req reqBody

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "Invalid request format"})
		return
	}

	defer r.Body.Close()

	if req.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "Name field is required"})
		return
	}

	varOpt := domain.NewVariationOption(req.Name, req.Value, parts[2])

	err = h.VariationSrv.CreateOption(context.TODO(), varOpt)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}
	w.WriteHeader(http.StatusCreated)
}
