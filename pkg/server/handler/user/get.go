package user

import (
	"encoding/json"
	"net/http"
)

type GetResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	HighScore int    `json:"highScore"`
	Coin      int    `json:"coin"`
}

func HandleGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			res GetResponse
		)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
