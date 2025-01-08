package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	config := setupServerConfig()
	mux := http.NewServeMux()

	styles := http.FileServer(http.Dir("./views/stylesheets"))
	mux.Handle("/styles/", http.StripPrefix("/styles/", styles))

	mux.HandleFunc("/", homePageHandler)

	mux.HandleFunc("/login", loginPageHandler)

	mux.HandleFunc("/signup", signupPageHandler)

	mux.HandleFunc("POST /api/login", config.loginHandler)
	mux.HandleFunc("POST /api/users", config.signupHandler)

	mux.HandleFunc("POST /api/validate-password", validatePasswordHandler)

	log.Println("Serving app on http://localhost:4321/")
	http.ListenAndServe(":4321", mux)
}
