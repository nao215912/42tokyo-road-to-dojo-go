package handler

import (
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

		WriteJson(w, res)
	}
}
