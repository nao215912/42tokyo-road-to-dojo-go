package handler

import (
	"encoding/json"
	"net/http"
)

type RankingListResponse struct {
	Ranks []struct {
		UserId   string `json:"userId"`
		UserName string `json:"userName"`
		Rank     int    `json:"rank"`
		Score    int    `json:"score"`
	} `json:"ranks"`
}

func HandleRankingList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			res RankingListResponse
		)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
