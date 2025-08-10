package handler

import (
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) ProductsPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"pkg/app/admin/api/web/templates/base.html",
		"pkg/app/admin/api/web/templates/products.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	data := struct {
		ApiBaseUrl string
	}{
		ApiBaseUrl: h.ApiBaseUrl,
	}
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
