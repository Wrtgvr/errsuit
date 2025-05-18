package errsuit

type ErrorResponse struct {
	ErrMsg string `json:"error"`
	Typ    string `json:"type"`
}
