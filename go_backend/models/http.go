package models

type ResponseModel struct {
	Status_Code int               `json:"status_code"`
	Content     map[string]string `json:"content"`
}

type ExceptionModel struct {
	Status_Code int    `json:"status_code"`
	Content     string `json:"detail"`
}

type JWTModel struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Message     string `json:"message"`
}

func NewResponseModel(status_code int, msg_en string, msg_zh string) *ResponseModel {
	return &ResponseModel{
		Status_Code: status_code,
		Content:     map[string]string{"msg_en": msg_en, "msg_zh": msg_zh},
	}
}

func NewExceptionModel(status_code int, detail string) *ExceptionModel {
	return &ExceptionModel{
		Status_Code: status_code,
		Content:     detail,
	}
}

func NewJWTModel(token string, token_type string, message string) *JWTModel {
	return &JWTModel{
		AccessToken: token,
		TokenType:   token_type,
		Message:     message,
	}
}
