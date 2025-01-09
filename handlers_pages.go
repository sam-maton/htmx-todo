package main

import (
	"fmt"
	"net/http"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }

	err := Config.applyMainLayout(w, r, "index.html", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func loginPageHandler(w http.ResponseWriter, r *http.Request) {

	err := Config.applyMainLayout(w, r, "./views/pages/login.html", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func signupPageHandler(w http.ResponseWriter, r *http.Request) {

	err := Config.applyMainLayout(w, r, "./views/pages/sign-up.html", nil)
	if err != nil {
		fmt.Println(err)
	}
}
