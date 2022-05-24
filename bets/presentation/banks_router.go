package presentation

import (
	"encoding/json"
	"my-bets/bets/application"
	"net/http"

	"github.com/go-chi/chi"
)

func BanksRouter(banksService *application.BanksService) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", createBank(banksService))
		r.Get("/{id}", getBank(banksService))
	}
}

func createBank(banksService *application.BanksService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var requestBody struct {
			InitialValue float64 `json:"initial_value,omitempty"`
		}

		if dErr := json.NewDecoder(r.Body).Decode(&requestBody); dErr != nil {
			responseWriter(w, http.StatusBadRequest, ErrorResponse{dErr.Error()})
			return
		}

		bank, err := banksService.CreateABank(requestBody.InitialValue)
		if err != nil {
			responseWriter(w, http.StatusInternalServerError, ErrorResponse{err.Error()})
			return
		}
		responseWriter(w, http.StatusOK, bank)
	}
}

func getBank(banksService *application.BanksService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		bank, err := banksService.GetABank(id)
		if err != nil {
			responseWriter(w, http.StatusInternalServerError, ErrorResponse{err.Error()})
			return
		}
		responseWriter(w, http.StatusOK, bank)
	}
}
