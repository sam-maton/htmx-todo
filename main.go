package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	config := setupServerConfig()
	mux := http.NewServeMux()

	styles := http.FileServer(http.Dir("./views/stylesheets"))
	mux.Handle("/styles/", http.StripPrefix("/styles/", styles))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		cookie, err := r.Cookie("htmx-auth")

		if err != nil {
			http.Redirect(w, r, "/signup", http.StatusFound)
			return
		}

		fmt.Println(cookie)

		err = applyMainLayout(w, r, "index.html")
		if err != nil {
			fmt.Println(err)
		}
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		err := applyMainLayout(w, r, "./views/pages/login.html")
		if err != nil {
			fmt.Println(err)
		}
	})

	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {

		err := applyMainLayout(w, r, "./views/pages/sign-up.html")
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

	mux.HandleFunc("POST /api/users", config.signupHandler)

	log.Println("Serving app on http://localhost:4321/")
	http.ListenAndServe(":4321", mux)
}
