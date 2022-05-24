package presentation

import (
	"encoding/json"
	"my-bets/bets/application"
	"net/http"

	"github.com/go-chi/chi"
)

func BetsRouter(betsService *application.BetsService) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", placeABet(betsService))
		r.Delete("/{id}", undoABet(betsService))
	}
}

func placeABet(betsService *application.BetsService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var betDTO application.PlaceBetDTO

		if dErr := json.NewDecoder(r.Body).Decode(&betDTO); dErr != nil {
			responseWriter(w, http.StatusBadRequest, ErrorResponse{dErr.Error()})
			return
		}

		bet, err := betsService.PlaceABet(betDTO)
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
