package res

import "blueberry_homework/internal/domain/entities"

// GetNames response structure
type GetNamesResponse struct {
	Message string              `json:"message"`
	Data    []entities.NameEntity `json:"data"`
}

// 성공 응답 structure
type SuccessResponse struct {
	Message string `json:"message"`
}

// 실패 response structure
type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
