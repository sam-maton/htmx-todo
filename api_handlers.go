package main

import (
	"log"
	"net/http"

	"github.com/sam-maton/htmx-todo/internal/auth"
	"github.com/sam-maton/htmx-todo/internal/database"
)

func (config serverConfig) signupHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("signup-email")
	password := r.FormValue("signup-password")
	confirm := r.FormValue("signup-password-confirm")

	if password != confirm {
		sendErrorToast(w, "The passwords did not match.")
		return
	}

	hashedPW, err := auth.HashPassword(password)
	if err != nil {
		sendErrorToast(w, "There was an error with the sign up.")
		log.Println(err)
		return
	}

	userParams := database.CreateUserParams{
		Email:          email,
		HashedPassword: hashedPW,
	}

	_, err = config.db.CreateUser(r.Context(), userParams)
	if err != nil {
		sendErrorToast(w, "There was an error with the sign up.")
		log.Println(err)
		return
	}

	w.Header().Add("Hx-Redirect", "/")
}
