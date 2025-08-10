package handler

import (
	"encoding/json"
	"net/http"
)

func (h Handler) SignInAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	msg := struct {
		Msg string `json:"msg"`
	}{
		Msg: "signInAdmin handler",
	}
	json.NewEncoder(w).Encode(msg)
}
