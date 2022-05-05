package handler

import (
	"42tokyo-road-to-dojo-go/pkg/database/dao"
	"42tokyo-road-to-dojo-go/pkg/http/middleware"
	"encoding/json"
	"net/http"
)

type UserUpdateRequest struct {
	Name string `json:"name"`
}

func HandleUserUpdate(d dao.Dao) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req UserUpdateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		user, err := middleware.UserOf(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		user.Name = req.Name

		user, err = d.NewUser().UpdateByName(r.Context(), user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
