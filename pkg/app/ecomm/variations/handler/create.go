package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/message"
	"github.com/amorindev/headless-ecomm-cms/pkg/app/ecomm/variations/domain"
)

func (h Handler) Create(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

    type reqBody struct {
		Name string `json:"name"`
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

    variation := domain.NewVariation(req.Name)

    err = h.VariationSrv.Create(context.TODO(), variation)
	if err != nil {
        w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
    }	
    
    w.WriteHeader(http.StatusCreated)
}

