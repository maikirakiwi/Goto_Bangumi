package api

import (
	"net/http"

	"github.com/rs/zerolog/log"
	json "github.com/sugawarayuuta/sonnet"

	db "Auto_Bangumi/v2/database"
	"Auto_Bangumi/v2/models/store"
)

func GetConfigHandler(w http.ResponseWriter, r *http.Request) {
	config, err := store.BoxForConfigModel(db.Conn).Get(1)
	if err != nil {
		log.Fatal().Msg("Error getting config.")
		writeResponse(w, r, 406, "Get config failed.", "获取配置失败。")
		return
	}

	resp, _ := json.Marshal(config)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}

func UpdateConfigHandler(w http.ResponseWriter, r *http.Request) {
	config := new(store.ConfigModel)
	config.Id = 1
	err := json.NewDecoder(r.Body).Decode(&config)
	if err != nil {
		log.Fatal().Msgf("Error decoding config: %s", err.Error())
		writeResponse(w, r, 406, "Update config failed.", "更新配置失败。")
		return
	}

	err = store.BoxForConfigModel(db.Conn).Update(config)
	if err != nil {
		log.Fatal().Msgf("Error updating config [objectBox]: %s", err.Error())
		writeResponse(w, r, 406, "Update config failed.", "更新配置失败。")
		return
	}

	writeResponse(w, r, 200, "Update config successfully.", "更新配置成功。")

}
