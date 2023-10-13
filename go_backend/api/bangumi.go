package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
	json "github.com/sugawarayuuta/sonnet"

	db "Auto_Bangumi/v2/database"
	dl "Auto_Bangumi/v2/downloaders"
	"Auto_Bangumi/v2/models/store"
)

func getAllBangumiHandler(w http.ResponseWriter, r *http.Request) {
	res, err := store.BoxForBangumi(db.Conn).GetAll()
	if err != nil {
		log.Error().Msgf("Error on /api/v1/bangumi/get/all: %s", err)
		writeException(w, r, 500, "Internal Server Error")
		return
	}

	json, err := json.Marshal(res)
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

	res, err := store.BoxForBangumi(db.Conn).Get(uint64(id))
	if err != nil || res == nil {
		log.Error().Msgf("Error on /api/v1/bangumi/get/{bangumi_id}: %s", err)
		writeResponse(w, r, 406, fmt.Sprintf("Can't find data with %d", id), fmt.Sprintf("无法找到 id %d 的数据", id))
		return
	}

	json, err := json.Marshal(res)
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
		log.Error().Msgf("Error on /api/v1/bangumi/update/{bangumi_id}: %d", err)
		writeResponse(w, r, 406, fmt.Sprintf("Can't find data with %d", id), fmt.Sprintf("无法找到 id %d 的数据", id))
		return
	}

	// Parse JSON from req body
	newData := new(store.Bangumi)
	newData.ID = uint64(id)
	err = json.NewDecoder(r.Body).Decode(&newData)
	if err != nil {
		log.Error().Msgf("Error on /api/v1/bangumi/update/{bangumi_id}: %d", err)
		writeResponse(w, r, 406, fmt.Sprintf("Can't find data with %d", id), fmt.Sprintf("无法找到 id %d 的数据", id))
		return
	}

	// Fetch existing Bangumi from database
	oldData, err := store.BoxForBangumi(db.Conn).Get(uint64(id))
	if err != nil {
		log.Error().Msgf("Error on /api/v1/bangumi/update/{bangumi_id}: %d", err)
		writeResponse(w, r, 406, fmt.Sprintf("Can't find data with %d", id), fmt.Sprintf("无法找到 id %d 的数据", id))
		return
	}

	// Generate new path
	newPath := fmtSavePath(oldData)
	if newPath == "" {
		log.Error().Msgf("Error on /api/v1/bangumi/update/{bangumi_id}: %s", err)
		writeException(w, r, 500, "Internal Server Error")
		return
	}

	// Check if path changed
	if oldData.SavePath != newPath {
		// Match torrents list
		list := dl.GetAllTorrent()
		if list == nil {
			log.Error().Msgf("Error on /api/v1/bangumi/update/{bangumi_id}: %s", err)
			writeException(w, r, 500, "Internal Server Error")
			return
		}

		changeQueue := []string{}
		for _, torrent := range list {
			if torrent.SavePath == oldData.SavePath {
				changeQueue = append(changeQueue, torrent.Hash)
			}
		}

		// Move torrent
		if len(changeQueue) > 0 {
			err = dl.Qbit.SetLocation(newPath, changeQueue...)
			if err != nil {
				log.Error().Msgf("Error on /api/v1/bangumi/update/{bangumi_id}: %s", err)
				writeException(w, r, 500, "Internal Server Error")
				return
			}
		}
	}

	// DB Transaction
	newData.SavePath = newPath
	err = store.BoxForBangumi(db.Conn).Update(newData)
	if err != nil {
		log.Error().Msgf("Error on /api/v1/bangumi/update/{bangumi_id}: %s", err)
		writeException(w, r, 500, "Internal Server Error")
		return
	}

	writeResponse(w, r, 200, "Update bangumi successfully.", "更新番剧成功。")
}

func fmtSavePath(b *store.Bangumi) string {
	var folder string
	if b.Year != "" {
		folder = fmt.Sprintf("%s (%s)", b.OfficialTitle, b.Year)
	} else {
		folder = b.OfficialTitle
	}

	cfg, err := store.BoxForConfigModel(db.Conn).Get(1)
	if err != nil {
		log.Fatal().Msg("Config DNE while formatting save path.")
		return ""
	}

	return fmt.Sprintf("%s/%s/Season %d", cfg.Downloader.Path, folder, b.Season)
}
