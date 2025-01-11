package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	mux := http.NewServeMux()
	config := setupServerConfig()

	styles := http.FileServer(http.Dir("./views/stylesheets"))
	mux.Handle("/styles/", http.StripPrefix("/styles/", styles))

	// PAGES
	mux.HandleFunc("/", config.middlewareAuth(config.homePageHandler))
	mux.HandleFunc("/login", config.loginPageHandler)
	mux.HandleFunc("/signup", config.signupPageHandler)

	// API AUTH
	mux.HandleFunc("POST /api/login", config.loginHandler)
	mux.HandleFunc("POST /api/signup", config.signupHandler)
	mux.HandleFunc("GET /api/logout", config.logoutHandler)

	// API TODOS
	mux.HandleFunc("POST /api/todos", config.middlewareAuth(config.createTodoHandler))
	mux.HandleFunc("DELETE /api/todos/{id}", config.middlewareAuth(config.deleteTodoHandler))
	mux.HandleFunc("PUT /api/todos/{id}", config.middlewareAuth(config.completedTodoHandler))

	// API VALIDATION
	mux.HandleFunc("POST /api/validate-password", validatePasswordHandler)

	log.Println("Serving app on http://localhost:4321/")
	http.ListenAndServe(":4321", mux)
}
