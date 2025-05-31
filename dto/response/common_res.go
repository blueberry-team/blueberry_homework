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

type TokenResponse struct {
	Token string `json:"token"`
}

type GetTokenResponse struct {
	Message string        `json:"message"`
	Data    TokenResponse `json:"data"`
}
