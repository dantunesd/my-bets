package presentation

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type HttpHandler struct {
	banksHandler *BanksHandler
	betsHandler  *BetsHandler
}

func NewHandler(banksHandler *BanksHandler, betsHandler *BetsHandler) *HttpHandler {
	return &HttpHandler{
		banksHandler: banksHandler,
		betsHandler:  betsHandler,
	}
}

func (h *HttpHandler) Create() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Post("/banks/", h.banksHandler.CreateBank())
	router.Get("/banks", h.banksHandler.GetBank())

	router.Post("/bets/", h.betsHandler.PlaceABet())
	router.Delete("/bets/{id}", h.betsHandler.UndoABet())

	return router
}
