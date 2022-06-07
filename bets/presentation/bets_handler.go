package presentation

import (
	"encoding/json"
	"my-bets/bets/application"
	"net/http"

	"github.com/go-chi/chi"
)

type BetsHandler struct {
	betsService *application.BetsService
}

func NewBetsHandler(betsService *application.BetsService) *BetsHandler {
	return &BetsHandler{
		betsService: betsService,
	}
}

func (h *BetsHandler) PlaceABet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var betDTO application.PlaceBetDTO

		if dErr := json.NewDecoder(r.Body).Decode(&betDTO); dErr != nil {
			responseWriter(w, http.StatusBadRequest, ErrorResponse{dErr.Error()})
			return
		}

		bet, err := h.betsService.PlaceABet(betDTO)
		if err != nil {
			responseWriter(w, http.StatusInternalServerError, ErrorResponse{err.Error()})
			return
		}
		responseWriter(w, http.StatusNoContent, bet)
	}
}

func (h *BetsHandler) UndoABet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if err := h.betsService.UndoABet(id); err != nil {
			responseWriter(w, http.StatusInternalServerError, ErrorResponse{err.Error()})
			return
		}
		responseWriter(w, http.StatusOK, nil)
	}
}
