package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/sam-maton/htmx-todo/internal/auth"
	"github.com/sam-maton/htmx-todo/internal/database"
)

func (config serverConfig) signupHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("signup-email")
	password := r.FormValue("signup-password")
	confirm := r.FormValue("signup-password-confirm")

	fmt.Println(email, password, confirm)

	// check if passwords match

	hashed, _ := auth.HashPassword(password)

	// create uuid

	userParams := database.CreateUserParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, _ := config.db.CreateUser(r.Context(), userParams)
}
