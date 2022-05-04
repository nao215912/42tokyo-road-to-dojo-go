package handler

import (
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

		WriteJson(w, &res)
	}
}
