package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	styles := http.FileServer(http.Dir("./views/stylesheets"))
	mux.Handle("/styles/", http.StripPrefix("/styles/", styles))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("").ParseFiles("index.html", "./views/layouts/main-layout.html")
		if err != nil {
			fmt.Println(err)
		}

		err = tmpl.ExecuteTemplate(w, "main-layout", nil)
		if err != nil {
			fmt.Println(err)
		}
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("").ParseFiles("./views/pages/login.html", "./views/layouts/main-layout.html")
		if err != nil {
			fmt.Println(err)
		}

		err = tmpl.ExecuteTemplate(w, "main-layout", nil)
		if err != nil {
			fmt.Println(err)
		}
	})

	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("").ParseFiles("./views/pages/sign-up.html", "./views/layouts/main-layout.html")
		if err != nil {
			fmt.Println(err)
		}

		err = tmpl.ExecuteTemplate(w, "main-layout", nil)
		if err != nil {
			fmt.Println(err)
		}
	})

	mux.HandleFunc("POST /api/login", func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("login-email")
		password := r.FormValue("login-password")

		fmt.Println(email)
		fmt.Println(password)
	})

	mux.HandleFunc("POST /api/validate-password", func(w http.ResponseWriter, r *http.Request) {
		password := r.FormValue("signup-password")
		confirm := r.FormValue("signup-password-confirm")

		type responseStruct struct {
			Value     string
			IsInvalid bool
		}

		data := responseStruct{
			Value:     confirm,
			IsInvalid: password != confirm,
		}

		view, err := template.ParseFiles("./views/components/confirm-password.html")

		if err != nil {
			log.Println(err)
		}

		view.Execute(w, data)
	})

	log.Println("Serving app on http://localhost:4321/")
	http.ListenAndServe(":4321", mux)
}
