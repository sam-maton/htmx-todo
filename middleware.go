package main

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/sam-maton/htmx-todo/internal/auth"
)

func (config serverConfig) middlewareAuth(handler func(http.ResponseWriter, *http.Request, uuid.UUID)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie(config.cookieTokenKey)
		if err != nil {
			log.Println("Redirecting to login page")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		userID, err := auth.ValidateJWT(cookie.Value, config.jwtSecret)
		if err != nil {
			log.Println("Redirecting to login page")
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		_, err = config.db.GetUserById(r.Context(), userID)
		if err != nil {
			log.Println("Redirecting to login page")
			cookie.MaxAge = -1
			cookie.Expires = time.Now()
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		handler(w, r, userID)
	}
}
