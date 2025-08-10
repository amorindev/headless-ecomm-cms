package api

import (
	"html/template"
	"net/http"
)

type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer() *TemplateRenderer {
	tmpl := template.Must(template.ParseGlob("pkg/app/admin/api/web/templates/*.html"))
	return &TemplateRenderer{templates: tmpl}

}

func (r TemplateRenderer) Render(w http.ResponseWriter, name string, data interface{}) {
	err := r.templates.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
