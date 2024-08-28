package v1

import (
	"encoding/json"
	"net/http"
)

func errorResponse(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}
