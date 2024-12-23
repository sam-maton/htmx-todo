package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("").ParseFiles("index.html", "./layouts/main-layout.html")
		if err != nil {
			fmt.Println(err)
		}

		err = tmpl.ExecuteTemplate(w, "main-layout", nil)
		if err != nil {
			fmt.Println(err)
		}
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("").ParseFiles("./pages/login.html", "./layouts/main-layout.html")
		if err != nil {
			fmt.Println(err)
		}

		err = tmpl.ExecuteTemplate(w, "main-layout", nil)
		if err != nil {
			fmt.Println(err)
		}
	})

	mux.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.New("").ParseFiles("./pages/sign-up.html", "./layouts/main-layout.html")
		if err != nil {
			fmt.Println(err)
		}

		err = tmpl.ExecuteTemplate(w, "main-layout", nil)
		if err != nil {
			fmt.Println(err)
		}
	})

	mux.HandleFunc("POST /api/login", func(w http.ResponseWriter, r *http.Request) {
		email := r.FormValue("login-email")
		password := r.FormValue("login-password")

		fmt.Println(email)
		fmt.Println(password)
	})

	mux.HandleFunc("POST /api/validate-password", func(w http.ResponseWriter, r *http.Request) {
		password := r.FormValue("signup-password")
		confirm := r.FormValue("signup-password-confirm")

		if password == confirm {
			io.WriteString(w, fmt.Sprintf(`
				<div hx-target="this" hx-swap="outerHTML">
					<label for="signup-password-confirm">Confirm Password:</label>
					<input
						class="valid-password"
						type="password"
						id="signup-password-confirm"
						name="signup-password-confirm"
						required
						value="%s"
						hx-post="/api/validate-password"
					/>
				</div>
			`, confirm))
		} else {
			io.WriteString(w, fmt.Sprintf(`
				<div hx-target="this" hx-swap="outerHTML">
					<label for="signup-password-confirm">Confirm Password:</label>
					<input
						class="invalid-password"
						type="password"
						id="signup-password-confirm"
						name="signup-password-confirm"
						required
						value="%s"
						hx-post="/api/validate-password"
					/>
					<p class="password-message--error">The two passwords do not match.</p>
				</div>
			`, confirm))
		}
	})

	log.Println("Serving app on http://localhost:4321/")
	http.ListenAndServe(":4321", mux)
}
