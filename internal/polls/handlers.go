package polls

import (
	"net/http"

	"github.com/asteroid2k/polln-api/internal/config"
	"github.com/asteroid2k/polln-api/internal/helpers"
	"github.com/go-chi/chi/v5"
)

func GetPolls(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := helpers.AppResponse{Status: 200, Data: map[string]any{"message": "HELLO"}}
		helpers.SendJSON(w, response)
	}
}

func RegisterRoutes(app *config.App, r chi.Router) {
	r.Group(func(r chi.Router) {
		r.Get("/polls", GetPolls(app))
	})

}
