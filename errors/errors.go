package errors

import "net/http"

type ErrorInfo struct {
	HttpStatus int       `json:"httpStatus"`
	ErrorBody  ErrorBody `json:"json:error"`
}

type ErrorBody struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

var BadRequestError = ErrorInfo{
	HttpStatus: http.StatusBadRequest,
	ErrorBody: ErrorBody{
		Code:    400,
		Status:  "Bad request",
		Message: "Bad request",
	},
}
