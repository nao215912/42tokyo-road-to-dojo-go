package handler

import (
	"encoding/json"
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

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(&res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
