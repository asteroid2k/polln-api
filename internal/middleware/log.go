package middleware

import (
	"net/http"

	"github.com/rs/zerolog/log"
)

func RequestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Info().Str("path", r.URL.Path).Str("method", r.Method).Str("user-agent", r.UserAgent()).Send()
			next.ServeHTTP(w, r)
		})
}
