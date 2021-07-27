package errors

import "net/http"

type ErrorInfo struct {
	HttpStatus int       `json:"httpStatus"`
	ErrorBody  ErrorBody `json:"error"`
}

func (e ErrorInfo) IsNil() bool {
	return ErrorInfo{} == e
}

type ErrorBody struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var NoError = ErrorInfo{
	HttpStatus: http.StatusOK,
	ErrorBody: ErrorBody{
		Code: 200,
		Msg:  "Successful",
	},
}

var BadRequestError = ErrorInfo{
	HttpStatus: http.StatusBadRequest,
	ErrorBody: ErrorBody{
		Code: 4000001,
		Msg:  "Invalid long url",
	},
}

var AliasForbidenError = ErrorInfo{
	HttpStatus: http.StatusForbidden,
	ErrorBody: ErrorBody{
		Code: 4030001,
		Msg:  "Invalid alias.",
	},
}

var UrlNotFoundError = ErrorInfo{
	HttpStatus: http.StatusNotFound,
	ErrorBody: ErrorBody{
		Code: 404,
		Msg:  "URL is not found.",
	},
}

var InternalServerError = ErrorInfo{
	HttpStatus: http.StatusInternalServerError,
	ErrorBody: ErrorBody{
		Code: 500,
		Msg:  "Internal Server Error",
	},
}
