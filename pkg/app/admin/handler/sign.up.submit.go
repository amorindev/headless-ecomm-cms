package handler

import (
	"encoding/json"
	"net/http"
)

func (h Handler) SignUpAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	msg := struct {
		Msg string `json:"msg"`
	}{
        Msg: "signUpAdmin handler",
    }
	json.NewEncoder(w).Encode(msg)
}
