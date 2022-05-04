package user

import (
	"encoding/json"
	"net/http"
)

type CreateRequest struct {
	Name string `json:"name"`
}

type CreateResponse struct {
	Token string `json:"token"`
}

func HandleCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req CreateRequest
			res CreateResponse
		)

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
