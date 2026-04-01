package model

type ErrorResponse struct {
	Code    int16  `json:"code"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}
