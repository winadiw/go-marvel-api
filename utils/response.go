package utils

type ResponseErrorData struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

func ResponseError(code int, message string, data interface{}) *ResponseErrorData {
	return &ResponseErrorData{
		Code:    code,
		Message: message,
		Status:  "error",
		Data:    data,
	}
}
