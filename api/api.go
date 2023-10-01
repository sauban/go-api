package api

import (
	"encoding/json"
	"net/http"
)

// Coin Balance params struct
type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	// Status code, usually 200
	Code int

	// Account balance
	Balance int64
}

type Error struct {
	// Status code, usually 500
	Code int

	// Error message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Unexpected error occurred.", http.StatusInternalServerError)
	}
)
