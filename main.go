package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var Config = setupServerConfig()

func main() {

	mux := http.NewServeMux()

	styles := http.FileServer(http.Dir("./views/stylesheets"))
	mux.Handle("/styles/", http.StripPrefix("/styles/", styles))

	// PAGES
	mux.HandleFunc("/", Config.middlewareAuth(homePageHandler))
	mux.HandleFunc("/login", loginPageHandler)
	mux.HandleFunc("/signup", signupPageHandler)

	// API AUTH
	mux.HandleFunc("POST /api/login", Config.loginHandler)
	mux.HandleFunc("POST /api/users", Config.signupHandler)
	mux.HandleFunc("GET /api/logout", Config.logoutHandler)

	// API VALIDATION
	mux.HandleFunc("POST /api/validate-password", validatePasswordHandler)

	log.Println("Serving app on http://localhost:4321/")
	http.ListenAndServe(":4321", mux)
}
