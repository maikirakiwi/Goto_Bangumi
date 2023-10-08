package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fatih/structs"
	"github.com/go-chi/chi/v5"
	"github.com/ostafen/clover/v2/query"
	"github.com/rs/zerolog/log"
	json "github.com/sugawarayuuta/sonnet"

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
		bangumi = append(bangumi, new(models.Bangumi).FromDocument(doc))
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
	id, err := fetchBangumiID(r)
	if err != nil {
		log.Error().Msgf("Error on /api/v1/bangumi/get/{bangumi_id}: %s", err)
		writeResponse(w, r, 406, fmt.Sprintf("Can't find data with %d", id), fmt.Sprintf("无法找到 id %d 的数据", id))
		return
	}

	res, err := db.FindOne("bangumi", "ID", id)
	if err != nil || res == nil {
		log.Error().Msgf("Error on /api/v1/bangumi/get/{bangumi_id}: %s", err)
		writeResponse(w, r, 406, fmt.Sprintf("Can't find data with %d", id), fmt.Sprintf("无法找到 id %d 的数据", id))
		return
	}

	json, err := json.Marshal(new(models.Bangumi).FromDocument(res))
	if err != nil {
		log.Error().Msgf("Error on /api/v1/bangumi/get/{bangumi_id}: %s", err)
		writeException(w, r, 500, "Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(json)

}

func fetchBangumiID(r *http.Request) (int, error) {
	return strconv.Atoi(chi.URLParam(r, "bangumi_id"))
}

func updateBangumiHandler(w http.ResponseWriter, r *http.Request) {
	id, err := fetchBangumiID(r)
	if err != nil {
		log.Error().Msgf("Error on /api/v1/bangumi/get/{bangumi_id}: %d", err)
		writeResponse(w, r, 406, fmt.Sprintf("Can't find data with %d", id), fmt.Sprintf("无法找到 id %d 的数据", id))
		return
	}
	body := models.BangumiUpdate{}
	json.NewDecoder(r.Body).Decode(&body)
	db.Conn.Update(query.NewQuery("bangumi").Where(query.Field("ID").Eq(id)), structs.Map(body))
	// TO DO: Implement save path update logic from bangumi.py

	writeResponse(w, r, 200, "Update bangumi successfully.", "更新番剧成功。")
}
