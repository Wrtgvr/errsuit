package errsuit

func BuildErrorResp(appErr *AppError) ErrorResponse {
	msg := appErr.Message
	if appErr.Type == TypeInternal {
		msg = "internal server error"
	}
	return ErrorResponse{
		msg,
		appErr.Type,
	}
}
