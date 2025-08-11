package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/amorindev/headless-ecomm-cms/cmd/api/message"
)

func (h Handler) DeleteVarOption(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/var-options/"):]

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: "missing category id"})
		return
	}

	err := h.VariationSrv.DeleteOption(context.TODO(), id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(message.ErrorMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
}
