package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func applyMainLayout(w http.ResponseWriter, r *http.Request, content string) error {

	path := r.URL.Path

	tmpl, err := template.New("").ParseFiles(content, "./views/layouts/main-layout.html")
	if err != nil {
		return fmt.Errorf("there was an error parsing the templates: %w", err)
	}

	err = tmpl.ExecuteTemplate(w, "main-layout", path)
	if err != nil {
		return fmt.Errorf("there was an error exectuing the temolate: %w", err)
	}

	return nil
}
