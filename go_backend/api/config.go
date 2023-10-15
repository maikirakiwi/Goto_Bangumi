package api

import (
	"net/http"

	"github.com/rs/zerolog/log"
	json "github.com/sugawarayuuta/sonnet"

	db "Auto_Bangumi/v2/database"
	"Auto_Bangumi/v2/models"
)

func GetConfigHandler(w http.ResponseWriter, r *http.Request) {
	config, exists := db.Cache.Get("config")
	if !exists {
		log.Fatal().Msg("Error getting config.")
	}

	resp, _ := json.Marshal(config.(models.Config))
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func UpdateConfigHandler(w http.ResponseWriter, r *http.Request) {
	var config models.Config
	err := json.NewDecoder(r.Body).Decode(&config)
	if err != nil {
		log.Fatal().Msgf("Error decoding config: %s", err.Error())
		writeResponse(w, r, 406, "Update config failed.", "更新配置失败。")
		return
	}

	err = db.Conn.Model(&models.Config{}).Where("ID = ?", 1).Save(&config).Error
	if err != nil {
		log.Fatal().Msgf("Error updating config: %s", err.Error())
		writeResponse(w, r, 406, "Update config failed.", "更新配置失败。")
		return
	}

	db.Cache.Set("config", config, -1)

	writeResponse(w, r, 200, "Update config successfully.", "更新配置成功。")

}
