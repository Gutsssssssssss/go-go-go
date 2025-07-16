package server

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		slog.Error("Error", "err", err)
	}
	if code > 499 {
		slog.Warn("Responding with 5XX", "msg", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
		Code  int    `json:"code"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
		Code:  code,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		slog.Error("marshalling JSON", "err", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	_, _ = w.Write(dat)
}
