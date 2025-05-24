package response

// 성공 응답 structure
type SuccessResponse struct {
	Message string `json:"message"`
}

// 실패 response structure
type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
