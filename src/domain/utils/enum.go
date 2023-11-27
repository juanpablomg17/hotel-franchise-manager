package utils

type GenericCommandResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
