package presentation

import (
	"my-bets/bets/application"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Router struct {
	banksService *application.BanksService
	betsService  *application.BetsService
}

func NewHandler(banksService *application.BanksService, betsService *application.BetsService) *Router {
	return &Router{
		banksService: banksService,
		betsService:  betsService,
	}
}

func (r *Router) Create() http.Handler {
	handler := chi.NewRouter()

	handler.Use(middleware.Logger)
	handler.Use(middleware.Recoverer)
	handler.Route("/banks", BanksRouter(r.banksService))
	handler.Route("/bets", BetsRouter(r.betsService))

	return handler
}
