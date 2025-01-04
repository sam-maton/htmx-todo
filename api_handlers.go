package main

import (
	"fmt"
	"net/http"
)

func (config serverConfig) signupHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("signup-email")
	password := r.FormValue("signup-password")
	confirm := r.FormValue("signup-password-confirm")

	fmt.Println(email, password, confirm)

	// check if passwords match

	// hash password

	// create uuid

	// userParams := database.CreateUserParams{

	// }

	// user, err := config.db.CreateUser(r.Context(), userParams)
}
