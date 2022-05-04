package handler

import (
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

		WriteJson(w, res)
	}
}
