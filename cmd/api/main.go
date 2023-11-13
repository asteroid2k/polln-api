package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/asteroid2k/polln-api/internal/config"
	"github.com/asteroid2k/polln-api/internal/users"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func main() {
	env := config.InitConfig()
	app := &config.App{}
	if env.ReleaseStage == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)
	mux.HandleFunc("/users/view", users.GetUser(app))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", env.Port),
		Handler: mux,
	}

	log.Info().Msg(fmt.Sprintf("server listening on %v", env.Port))
	err := srv.ListenAndServe()
	log.Fatal().Err(err).Msg("server could not listen")
}
