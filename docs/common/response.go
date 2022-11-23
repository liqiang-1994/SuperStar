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
