package templates

import (
	"bytes"
	"fmt"
	"html/template"
)

func LoadTemplate(templatePath string, data any) (string, error) {
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", fmt.Errorf("parse files - loadTemplate: %w", err)
	}

	var body bytes.Buffer

	err = t.Execute(&body, data)

	if err != nil {
		return "", fmt.Errorf("execute - load template err: %w", err)
	}

	return body.String(), nil
}
