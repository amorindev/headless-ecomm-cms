package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/message"
)

func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[len("/categories/"):]

    if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "missing category id"})
		return
	}

    err := h.VariationSrv.Delete(context.TODO(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}
