package handler

import (
	"42tokyo-road-to-dojo-go/pkg/http/middleware"
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
		// Todo: error handle
		res.Name, _ = middleware.UserOf(r)

		WriteJson(w, &res)
	}
}
