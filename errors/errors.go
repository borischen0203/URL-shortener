package errors

// type ErrorInfo struct {
// 	// HttpStatus int       `json:"httpStatus"`
// 	ErrorBody ErrorBody `json:"error"`
// }

type ErrorInfo struct {
	Code uint16 `json:"code"`
	Msg  string `json:"msg"`
}

func (e ErrorInfo) IsNil() bool {
	return ErrorInfo{} == e
}

var NoError = ErrorInfo{}

var InvalidLongUrlError = ErrorInfo{Code: 40001, Msg: "Invalid long url"}

var InvalidAliasError = ErrorInfo{
	Code: 40002,
	Msg:  "Invalid alias",
}

var AliasForbidenError = ErrorInfo{
	Code: 40301,
	Msg:  "Alias is used",
}

var UrlNotFoundError = ErrorInfo{
	Code: 404,
	Msg:  "URL is not found.",
}

var InternalServerError = ErrorInfo{
	Code: 500,
	Msg:  "Internal Server Error",
}
