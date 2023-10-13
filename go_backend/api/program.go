package api

import (
	"net/http"

	json "github.com/sugawarayuuta/sonnet"

	db "Auto_Bangumi/v2/database"
	"Auto_Bangumi/v2/models"
)

func ProgramStatusHandler(w http.ResponseWriter, r *http.Request) {
	cfg, exists := db.Cache.Get("config")
	res, err := json.Marshal(map[string]interface{}{
		"status":    exists,
		"version":   cfg.(models.ConfigModel).Program.DataVersion,
		"first_run": false,
	})
	if err != nil {
		writeException(w, r, 500, "Internal Server Error")
	} else {
		w.Write(res)
	}
}
