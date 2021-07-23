package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/draios/shared-go/pkg/sdauth"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog"

	pkgmiddleware "github.com/draios/secure-backend/pkg/http/middleware"
)

const (
	// API server listens on this port
	containerPort = 12345
	// API route begins with this prefix
	handlerPath = "/api/docs"
)

func main() {
	// Setup logger
	logger := zerolog.New(os.Stderr).With().
		Timestamp().
		Logger().
		Level(zerolog.InfoLevel)
	if os.Getenv("DEBUG") != "" {
		logger = logger.
			Output(zerolog.ConsoleWriter{
				Out: os.Stdout,
			}).
			Level(zerolog.DebugLevel)
	}

	endpoint := os.Getenv("MONITOR_BACKEND_ENDPOINT")
	if endpoint == "" {
		logger.Fatal().Msg("MONITOR_BACKEND_ENDPOINT not set")
	}

	authClient, err := getAuthClient(endpoint)
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}

	router := chi.NewRouter()
	router.Use(pkgmiddleware.Recoverer)
	router.Use(pkgmiddleware.Authentication(authClient))

	router.Route(handlerPath, func(r chi.Router) {
		r.Use(pkgmiddleware.Authorization(authClient, sdauth.Permission("secure.onboarding.read")))

		r.Get("/secure", apiDocHandler)
		r.Get("/secure/internal", apiInternalDocHandler)
	})

	logger.Info().Msg("Starting API documentation service")
	err = http.ListenAndServe(fmt.Sprintf(":%d", containerPort), router)
	if err != nil {
		logger.Error().Err(err)
		os.Exit(1)
	}
}

func apiDocHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/static/index.html")
}

func apiInternalDocHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/static/internal.html")
}
