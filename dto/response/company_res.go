package response

import (
	"time"

)

type CompanyResponse struct {
	CompanyName    string     `json:"companyName"`
	CompanyAddress string     `json:"companyAddress"`
	TotalStaff     int        `json:"totalStaff"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
}

// GetCompanies response structure
type GetCompaniesResponse struct {
	Message string                   `json:"message"`
	Data    CompanyResponse `json:"data"`
}
