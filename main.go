package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./pages/login.html")
	})

	log.Println("Serving app on http://localhost:4321/")
	http.ListenAndServe(":4321", mux)
}
