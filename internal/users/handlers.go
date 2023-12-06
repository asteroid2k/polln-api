package users

import (
	"net/http"

	"github.com/asteroid2k/polln-api/api"
	"github.com/asteroid2k/polln-api/internal/config"
	"github.com/asteroid2k/polln-api/internal/helpers"
	"github.com/go-chi/chi/v5"
)

func GetUsers(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Users"))
	}
}

func CreateUser(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userData := api.CreateUserRequest{}
		ok := helpers.ParseJSON(w, r.Body, &userData)
		if !ok {
			return
		}

		helpers.SendJSON(w, helpers.AppResponse{Data: userData})
	}
}

func RegisterRoutes(app *config.App, r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Post("/users", CreateUser(app))
		r.Get("/users", GetUsers(app))
	})

}
