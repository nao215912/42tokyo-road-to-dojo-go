package handler

import (
	"net/http"
)

type CollectionListResponse struct {
	Collections []struct {
		CollectionID string `json:"collectionID"`
		Name         string `json:"name"`
		Rarity       int    `json:"rarity"`
		HasItem      bool   `json:"hasItem"`
	} `json:"collections"`
}

func HandleCollectionList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			res CollectionListResponse
		)

		WriteJson(w, &res)
	}
}
