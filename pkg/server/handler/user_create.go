package handler

import (
	"42tokyo-road-to-dojo-go/pkg/database/dao"
	"encoding/json"
	"net/http"
)

type UserCreateRequest struct {
	Name string `json:"name"`
}

type UserCreateResponse struct {
	Token string `json:"token"`
}

func HandleUserCreate(d dao.Dao) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req UserCreateRequest
			res UserCreateResponse
		)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user, err := d.NewUser().Create(r.Context(), req.Name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		res.Token = user.Token

		WriteJson(w, &res)
	}
}
