package api

import (
	"encoding/json"
	"net/http"

	"github.com/ostafen/clover/v2/query"
	"github.com/rs/zerolog/log"

	db "Auto_Bangumi/v2/database"
	"Auto_Bangumi/v2/models"
)

func getAllBangumiHandler(w http.ResponseWriter, r *http.Request) {
	res, err := db.Conn.FindAll(query.NewQuery("bangumi"))
	if err != nil {
		log.Error().Msgf("Error on /api/v1/bangumi/get/all: %s", err)
		writeException(w, r, 500, "Internal Server Error")
	}

	// Encode list of documents to JSON using Bangumi struct
	bangumi := []models.Bangumi{}
	for _, doc := range res {
		bangumi = append(bangumi, *new(models.Bangumi).FromDocument(doc))
	}

	json, err := json.Marshal(bangumi)
	if err != nil {
		log.Error().Msgf("Error on /api/v1/bangumi/get/all: %s", err)
		writeException(w, r, 500, "Internal Server Error")
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(json)
}
