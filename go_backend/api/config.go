package api

import (
	"encoding/json"
	"net/http"

	"github.com/rs/zerolog/log"

	db "Auto_Bangumi/v2/database"
)

func GetConfigHandler(w http.ResponseWriter, r *http.Request) {
	config, err := db.FindOne("config", "backend", "go")
	if err != nil {
		log.Fatal().Msgf("Error getting config: %s", err.Error())
	}

	resp, _ := json.Marshal(config.Get("content"))
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func UpdateConfigHandler(w http.ResponseWriter, r *http.Request) {
	var config map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&config)
	if err != nil {
		log.Fatal().Msgf("Error decoding config: %s", err.Error())
		writeResponse(w, r, 406, "Update config failed.", "更新配置失败。")
		return
	}

	err = db.UpdateOne("config", "backend", "go", "content", config)
	if err != nil {
		log.Fatal().Msgf("Error updating config [Clover]: %s", err.Error())
		writeResponse(w, r, 406, "Update config failed.", "更新配置失败。")
		return
	}

	db.Cache.Set("config", config, -1)

	writeResponse(w, r, 200, "Update config successfully.", "更新配置成功。")

}
