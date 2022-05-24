package presentation

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func responseWriter(w http.ResponseWriter, code int, content interface{}) {
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(code)
	if content != nil {
		body, _ := json.Marshal(content)
		w.Write(body)
	}
}
