package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

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

	user, err := config.db.CreateUser(r.Context(), userParams)
	if err != nil {
		sendErrorToast(w, "There was an error with the sign up.")
		log.Println(err)
		return
	}

	token, _ := auth.MakeJWT(user.ID, config.jwtSecret)

	cookie := http.Cookie{
		Path:    "/",
		Name:    config.cookieTokenKey,
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
	}

	http.SetCookie(w, &cookie)

	w.Header().Add("Hx-Redirect", "/")
}

func (config serverConfig) loginHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("login-email")
	password := r.FormValue("login-password")

	user, err := config.db.GetUserByEmail(r.Context(), email)

	if err != nil {
		sendErrorToast(w, "No user with that email exists.")
		log.Println(err)
		return
	}

	err = auth.CheckPasswordHash(password, user.HashedPassword)
	if err != nil {
		sendErrorToast(w, "The password is incorrect.")
		log.Println(err)
		return
	}

	token, _ := auth.MakeJWT(user.ID, config.jwtSecret)

	cookie := http.Cookie{
		Path:    "/",
		Name:    config.cookieTokenKey,
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
	}

	http.SetCookie(w, &cookie)

	w.Header().Add("Hx-Redirect", "/")
}

func (config serverConfig) logoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(config.cookieTokenKey)
	if err != nil {
		log.Println("Redirecting to login page")
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	deleteCookie := http.Cookie{
		Path:   "/",
		Name:   cookie.Name,
		MaxAge: -1,
	}

	log.Println("Redirecting to login page")
	http.SetCookie(w, &deleteCookie)
	http.Redirect(w, r, "/login", http.StatusFound)
}

func validatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	password := r.FormValue("signup-password")
	confirm := r.FormValue("signup-password-confirm")

	type responseStruct struct {
		Value     string
		IsInvalid bool
	}

	data := responseStruct{
		Value:     confirm,
		IsInvalid: password != confirm,
	}

	view, err := template.ParseFiles("./views/components/confirm-password.html")

	if err != nil {
		log.Println(err)
	}

	view.Execute(w, data)
}
