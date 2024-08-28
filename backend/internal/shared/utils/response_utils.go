package utils

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON JSON応答を返します
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// RespondWithError エラー応答を返します
func RespondWithError(w http.ResponseWriter, code int, message string) {
	http.Error(w, message, code)
}
