package users

import (
	"net/http"

	"github.com/asteroid2k/polln-api/internal/config"
)

func GetUser(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("A User"))
	}
}
