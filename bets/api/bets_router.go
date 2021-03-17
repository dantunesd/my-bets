package api

import (
	"encoding/json"
	"my-bets/bets/application"
	"my-bets/bets/domain"
	"net/http"

	"github.com/go-chi/chi"
)

func placeABet(betsService *application.BetsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var bet domain.Bet

		if dErr := json.NewDecoder(r.Body).Decode(&bet); dErr != nil {
			responseWriter(w, http.StatusBadRequest, ErrorResponse{dErr.Error()})
			return
		}

		bet, err := betsService.PlaceABet(bet)
		if err != nil {
			responseWriter(w, http.StatusInternalServerError, ErrorResponse{err.Error()})
			return
		}
		responseWriter(w, http.StatusNoContent, bet)
	}
}

func undoABet(betsService *application.BetsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if err := betsService.UndoABet(id); err != nil {
			responseWriter(w, http.StatusInternalServerError, ErrorResponse{err.Error()})
			return
		}
		responseWriter(w, http.StatusOK, nil)
	}
}
func BetsRouter(betsService *application.BetsService) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", placeABet(betsService))
		r.Delete("/{id}", undoABet(betsService))
	}
}
