package handler

import (
	"encoding/json"
	"net/http"
)

func WriteJson(w http.ResponseWriter, v interface{}) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
