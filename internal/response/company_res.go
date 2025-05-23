package response

import "blueberry_homework/internal/domain/entities"

// GetCompanies response structure
type GetCompaniesResponse struct {
	Message string                   `json:"message"`
	Data    []entities.CompanyEntity `json:"data"`
}
