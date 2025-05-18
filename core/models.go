package errsuit

type ErrorResponse struct {
	ErrMsg string `json:"error"`
	Typ    string `json:"type"`
}

type ErrorHandler interface {
	HandleError(Context, error) bool
}

type Context interface {
	SetHeader(key, value string)
	GetHeader(key string) string
	SetStatus(code int)
	Write(body []byte)
}
