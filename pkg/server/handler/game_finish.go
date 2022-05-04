package handler

import (
	"encoding/json"
	"net/http"
)

type GameFinishResponse struct {
	Coin int `json:"coin"`
}

func HandleGameFinish() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			res GameFinishResponse
		)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
