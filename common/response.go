package common

import "net/http"

type ResponseModel struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  int         `json:"status"`
}

func Fail(err error) *ResponseModel {
	return &ResponseModel{
		Status:  http.StatusInternalServerError,
		Message: err.Error(),
	}
}

func Success(data interface{}) *ResponseModel {
	return &ResponseModel{
		Success: true,
		Data:    data,
		Message: "Success",
		Status:  http.StatusOK,
	}
}

func NotLogin(message string) *ResponseModel {
	return &ResponseModel{
		Success: false,
		Data:    nil,
		Message: message,
		Status:  http.StatusUnauthorized,
	}
}
