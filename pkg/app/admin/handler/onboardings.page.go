package handler

import (
	"html/template"
	"log"
	"net/http"
)

func (h *Handler) OnboardingsPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"pkg/app/admin/api/web/templates/base.html",
		"pkg/app/admin/api/web/templates/onboardings.html",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
