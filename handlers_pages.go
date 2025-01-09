package main

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/sam-maton/htmx-todo/internal/database"
)

func (config *serverConfig) homePageHandler(w http.ResponseWriter, r *http.Request, userId uuid.UUID) {
	// if r.URL.Path != "/" {
	// 	http.NotFound(w, r)
	// 	return
	// }
	todos, _ := config.db.GetTodosByUserId(r.Context(), userId)

	err := config.applyMainLayout(w, r, "./views/pages/index.html", struct {
		Todos []database.GetTodosByUserIdRow
	}{
		Todos: todos,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func (config *serverConfig) loginPageHandler(w http.ResponseWriter, r *http.Request) {

	err := config.applyMainLayout(w, r, "./views/pages/login.html", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func (config *serverConfig) signupPageHandler(w http.ResponseWriter, r *http.Request) {

	err := config.applyMainLayout(w, r, "./views/pages/sign-up.html", nil)
	if err != nil {
		fmt.Println(err)
	}
}
