package main

import (
	"fmt"
	"net/http"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	err := applyMainLayout(w, r, "index.html")
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
