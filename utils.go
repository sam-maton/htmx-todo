package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type MainLayoutData struct {
	IsAuthenticated bool
	ContentData     interface{}
}

func (config *serverConfig) applyMainLayout(w http.ResponseWriter, r *http.Request, content string, contentData interface{}) error {

	_, cookieErr := r.Cookie(config.cookieTokenKey)

	tmpl, err := template.New("").ParseFiles(content, "./views/layouts/main-layout.html")
	if err != nil {
		return fmt.Errorf("there was an error parsing the templates: %w", err)
	}

	data := MainLayoutData{
		IsAuthenticated: cookieErr == nil,
		ContentData:     contentData,
	}

	err = tmpl.ExecuteTemplate(w, "main-layout", data)
	if err != nil {
		return fmt.Errorf("there was an error exectuing the template: %w", err)
	}

	return nil
}

func sendErrorToast(w http.ResponseWriter, message string) {

	type errorParams struct {
		ErrorMessage string
	}

	view, err := template.ParseFiles("./views/components/error-toast.html")

	if err != nil {
		log.Println(err)
	}

	// w.WriteHeader(http.StatusInternalServerError)
	w.Header().Add("HX-Retarget", "#main-toast-container")
	w.Header().Add("HX-Reswap", "innerHTML")
	view.Execute(w, errorParams{ErrorMessage: message})
}
