package users

import (
	"net/http"

	"github.com/asteroid2k/polln-api/api"
	"github.com/asteroid2k/polln-api/internal/config"
	"github.com/asteroid2k/polln-api/internal/utils/helpers"
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
		vErrors, ok := app.Validator.ValidateStruct(userData)
		if !ok {
			helpers.SendJSON(w, helpers.NewValidationErrorResponse(vErrors, nil))
			return
		}

		helpers.SendJSON(w, helpers.AppResponse{Data: userData})
	}
}

func RegisterRoutes(app *config.App, r chi.Router) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", CreateUser(app))
		r.Get("/", GetUsers(app))
	})

}
