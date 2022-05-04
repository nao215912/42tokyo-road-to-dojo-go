package handler

import (
	"net/http"
)

type UserGetResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	HighScore int    `json:"highScore"`
	Coin      int    `json:"coin"`
}

func HandleUserGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			res UserGetResponse
		)

		WriteJson(w, res)
	}
}
