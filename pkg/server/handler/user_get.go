package handler

import (
	"42tokyo-road-to-dojo-go/pkg/http/middleware"
	"net/http"
)

type UserGetResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	HighScore int64  `json:"highScore"`
	Coin      int64  `json:"coin"`
}

func HandleUserGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := middleware.UserOf(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		res := UserGetResponse{
			Id:        user.ID,
			Name:      user.Name,
			HighScore: user.HighScore,
			Coin:      user.Coin,
		}
		WriteJson(w, &res)
	}
}
