package api

import (
	"bufio"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
	json "github.com/sugawarayuuta/sonnet"

	db "Auto_Bangumi/v2/database"
	"Auto_Bangumi/v2/models"
)

func ProgramStatusHandler(w http.ResponseWriter, r *http.Request) {
	cfg, exists := db.Cache.Get("config")
	res, err := json.Marshal(map[string]interface{}{
		"status":    exists,
		"version":   cfg.(models.Config).Program.DataVersion,
		"first_run": false,
	})
	if err != nil {
		writeException(w, r, 500, "Internal Server Error")
	} else {
		w.Write(res)
	}
}

func LogOutputHandler(w http.ResponseWriter, r *http.Request) {
	// read stderr and write to response as plain text
	file, err := os.OpenFile("log.txt", os.O_RDONLY, 0666)
	if err != nil {
		writeException(w, r, 500, "Internal Server Error")
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		w.Write([]byte(scanner.Text() + "\n"))
	}
}

func LogClearHandler(w http.ResponseWriter, r *http.Request) {
	ClearLog()

	writeResponse(w, r, 200, "Log cleared successfully.", "日志清除成功。")
}

func ClearLog() {
	// clear log.txt
	file, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	defer file.Close()
}
