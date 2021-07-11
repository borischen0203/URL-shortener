package errors

import "net/http"

type ErrorInfo struct {
	HttpStatus int       `json:"httpStatus"`
	ErrorBody  ErrorBody `json:"json:error"`
}

type ErrorBody struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var BadRequestError = ErrorInfo{
	HttpStatus: http.StatusBadRequest,
	ErrorBody: ErrorBody{
		Code:    400,
		Message: "Bad request",
	},
}

var AliasForbidenError = ErrorInfo{
	HttpStatus: http.StatusForbidden,
	ErrorBody: ErrorBody{
		Code:    403,
		Message: "Invalid alias.",
	},
}

var UrlNotFoundError = ErrorInfo{
	HttpStatus: http.StatusNotFound,
	ErrorBody: ErrorBody{
		Code:    404,
		Message: "URL is not found.",
	},
}

var InternalServerError = ErrorInfo{
	HttpStatus: http.StatusInternalServerError,
	ErrorBody: ErrorBody{
		Code:    500,
		Message: "Internal Server Error",
	},
}
