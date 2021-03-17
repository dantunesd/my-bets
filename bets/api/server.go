package api

import (
	"my-bets/bets/application"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func CreateAndStartServer(banksService *application.BanksService, betsService *application.BetsService) {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Route("/banks", BanksRouter(banksService))
	r.Route("/bets", BetsRouter(betsService))

	http.ListenAndServe(":8080", r)
}
