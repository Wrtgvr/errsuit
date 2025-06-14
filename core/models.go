package errsuit

type ErrorResponse struct {
	ErrMsg string `json:"error"`
	Typ    string `json:"type"`
	Code   int    `json:"code,omitempty"`
}

type ErrorHandler interface {
	HandleError(Context, error, ResponseFormat) bool
}

type Context interface {
	SetHeader(key, value string)
	GetHeader(key string) string
	SetStatus(code int)
	Write(body []byte)
}
