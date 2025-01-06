package main

import (
	"fmt"
	"net/http"

	"github.com/sam-maton/htmx-todo/internal/auth"
	"github.com/sam-maton/htmx-todo/internal/database"
)

func (config serverConfig) signupHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("signup-email")
	password := r.FormValue("signup-password")
	confirm := r.FormValue("signup-password-confirm")

	fmt.Println(email, password, confirm)

	if true {
		sendErrorToast(w)
		return
	}

	hashedPW, _ := auth.HashPassword(password)

	userParams := database.CreateUserParams{
		Email:          email,
		HashedPassword: hashedPW,
	}

	config.db.CreateUser(r.Context(), userParams)
}
