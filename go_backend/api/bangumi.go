package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
		return
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
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(json)
}

func getBangumiHandler(w http.ResponseWriter, r *http.Request) {
	bangumiID := chi.URLParam(r, "bangumi_id")
	bangumiINT, err := strconv.Atoi(bangumiID)
	if err != nil {
		log.Error().Msgf("Error on /api/v1/bangumi/get/{bangumi_id}: %s", err)
		writeResponse(w, r, 406, fmt.Sprintf("Can't find data with %s", bangumiID), fmt.Sprintf("无法找到 id %s 的数据", bangumiID))
		return
	}

	res, err := db.FindOne("bangumi", "ID", bangumiINT)
	if err != nil || res == nil {
		log.Error().Msgf("Error on /api/v1/bangumi/get/{bangumi_id}: %s", err)
		writeResponse(w, r, 406, fmt.Sprintf("Can't find data with %s", bangumiID), fmt.Sprintf("无法找到 id %s 的数据", bangumiID))
		return
	}

	json, err := json.Marshal(*new(models.Bangumi).FromDocument(res))
	if err != nil {
		log.Error().Msgf("Error on /api/v1/bangumi/get/{bangumi_id}: %s", err)
		writeException(w, r, 500, "Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(json)

}
