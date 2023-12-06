package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/asteroid2k/polln-api/internal/config"
	"github.com/asteroid2k/polln-api/internal/utils/helpers"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	env := config.ParseEnv()
	val, trans := helpers.NewValidator()

	app := &config.App{Env: env, Validator: helpers.AppValidator{Validate: val, Translator: trans}}

	if env.ReleaseStage == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", env.Port),
		Handler: routes(app),
	}

	log.Info().Msg(fmt.Sprintf("server listening on %v", env.Port))
	err := srv.ListenAndServe()
	log.Fatal().Err(err).Msg("server could not listen")
}
