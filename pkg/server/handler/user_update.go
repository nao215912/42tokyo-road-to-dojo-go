package handler

import (
	"encoding/json"
	"net/http"
)

type UserUpdateResponse struct {
	Name string `json:"name"`
}

func HandleUserUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			res UserUpdateResponse
		)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
