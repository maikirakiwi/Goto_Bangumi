package models

type ResponseModel struct {
	Status_Code int               `json:"status_code"`
	Content     map[string]string `json:"content"`
}

func NewResponseModel(status_code int, msg_en string, msg_zh string) *ResponseModel {
	return &ResponseModel{
		Status_Code: status_code,
		Content:     map[string]string{"msg_en": msg_en, "msg_zh": msg_zh},
	}
}
