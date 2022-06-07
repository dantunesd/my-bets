package presentation

import (
	"encoding/json"
	"my-bets/bets/application"
	"net/http"

	"github.com/go-chi/chi"
)

type BanksHandler struct {
	banksService *application.BanksService
}

func NewBanksHandler(banksService *application.BanksService) *BanksHandler {
	return &BanksHandler{
		banksService: banksService,
	}
}

func (h *BanksHandler) CreateBank() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var requestBody struct {
			InitialValue float64 `json:"initial_value,omitempty"`
		}

		if dErr := json.NewDecoder(r.Body).Decode(&requestBody); dErr != nil {
			responseWriter(w, http.StatusBadRequest, ErrorResponse{dErr.Error()})
			return
		}

		bank, err := h.banksService.CreateABank(requestBody.InitialValue)
		if err != nil {
			responseWriter(w, http.StatusInternalServerError, ErrorResponse{err.Error()})
			return
		}
		responseWriter(w, http.StatusOK, bank)
	}
}

func (h *BanksHandler) GetBank() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		bank, err := h.banksService.GetABank(id)
		if err != nil {
			responseWriter(w, http.StatusInternalServerError, ErrorResponse{err.Error()})
			return
		}
		responseWriter(w, http.StatusOK, bank)
	}
}
