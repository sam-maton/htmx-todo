package main

import (
	"fmt"
	"log"
	"net/http"
)

func (cfg serverConfig) middlewareAuth(handler func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(cfg.cookieTokenKey)

		if err != nil {
			log.Println("Redirecting to login page")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		fmt.Println(cookie)

		handler(w, r)
	}
}
