package handler

import (
	"net/http"
)

type GachaDrawResponse struct {
	Ranks []struct {
		UserId   string `json:"userId"`
		UserName string `json:"userName"`
		Rank     int    `json:"rank"`
		Score    int    `json:"score"`
	} `json:"ranks"`
}

func HandleGachaDraw() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			res GachaDrawResponse
		)

		WriteJson(w, &res)
	}
}
