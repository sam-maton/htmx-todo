package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// http.ServeFile(w, r, "index.html")
		tmpl, err := template.New("").ParseFiles("index.html", "./layouts/main-layout.html")
		if err != nil {
			fmt.Println(err)
		}
		// check your err
		err = tmpl.ExecuteTemplate(w, "main-layout", nil)

		if err != nil {
			fmt.Println(err)
		}
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// http.ServeFile(w, r, "./pages/login.html")
		tmpl, err := template.New("").ParseFiles("./pages/login.html", "./layouts/main-layout.html")
		if err != nil {
			fmt.Println(err)
		}
		// check your err
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

	mux.HandleFunc("GET /api/clicked", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<h1>HTMX is awesome!</h1>")
	})

	log.Println("Serving app on http://localhost:4321/")
	http.ListenAndServe(":4321", mux)
}
