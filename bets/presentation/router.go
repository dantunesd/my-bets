package presentation

import (
	"my-bets/bets/application"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func HandlersFactory(banksService *application.BanksService, betsService *application.BetsService) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Route("/banks", BanksRouter(banksService))
	r.Route("/bets", BetsRouter(betsService))

	return r
}
