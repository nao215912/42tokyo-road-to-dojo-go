package user

import (
	"encoding/json"
	"net/http"
)

type UpdateResponse struct {
	Name string `json:"name"`
}

func HandleUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			res UpdateResponse
		)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
