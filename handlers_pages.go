package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	cookie, err := r.Cookie("htmx-auth")

	if err != nil {
		log.Println("Redirecting to login page")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	fmt.Println(cookie)

	err = applyMainLayout(w, r, "index.html")
	if err != nil {
		fmt.Println(err)
	}
}

func loginPageHandler(w http.ResponseWriter, r *http.Request) {

	err := applyMainLayout(w, r, "./views/pages/login.html")
	if err != nil {
		fmt.Println(err)
	}
}

func signupPageHandler(w http.ResponseWriter, r *http.Request) {

	err := applyMainLayout(w, r, "./views/pages/sign-up.html")
	if err != nil {
		fmt.Println(err)
	}
}
