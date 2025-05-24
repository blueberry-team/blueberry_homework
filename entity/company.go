package entity

import (
	"time"

	"github.com/google/uuid"
)

// CompanyEntity 회사 정보를 나타내는 구조체
type CompanyEntity struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`        // 사용자 이름
	CompanyName string    `json:"company_name"`
	CreatedAt   time.Time `json:"created_at"`
}

// NewCompanyEntity 새로운 CompanyEntity 생성
func NewCompanyEntity(name, companyName string) CompanyEntity {
	return CompanyEntity{
		ID:          uuid.New().String(),
		Name:        name,
		CompanyName: companyName,
		CreatedAt:   time.Now(),
	}
}
