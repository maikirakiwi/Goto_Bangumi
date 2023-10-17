package api

import (
	"net/http"

	json "github.com/sugawarayuuta/sonnet"

	models "Auto_Bangumi/v2/models"
)

func writeResponse(w http.ResponseWriter, r *http.Request, statusCode int, msg_en string, msg_zh string) {
	response, _ := json.Marshal(models.NewResponseModel(statusCode, msg_en, msg_zh))
	w.WriteHeader(statusCode)
	w.Write(response)
}

func writeJWTResponse(w http.ResponseWriter, r *http.Request, token string, token_type string, message string) {
	response, _ := json.Marshal(models.NewJWTModel(token, token_type, message))
	w.Write(response)
}

func writeException(w http.ResponseWriter, r *http.Request, statusCode int, detail string) {
	response, _ := json.Marshal(models.NewExceptionModel(statusCode, detail))
	w.WriteHeader(statusCode)
	w.Write(response)
}
