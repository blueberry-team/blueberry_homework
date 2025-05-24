package entity

import (
	"time"

	"github.com/google/uuid"
)

// CompanyEntity 회사 정보를 나타내는 구조체 (확장된 버전)
type CompanyEntity struct {
	ID             string    `json:"id" bson:"id"`
	UserID         string    `json:"userId" bson:"userId"`               // 회사 소유자 ID
	CompanyName    string    `json:"companyName" bson:"companyName"`
	CompanyAddress string    `json:"companyAddress" bson:"companyAddress"`
	TotalStaff     int       `json:"totalStaff" bson:"totalStaff"`
	CreatedAt      time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt" bson:"updatedAt"`
}

// NewCompanyEntity 새로운 CompanyEntity 생성
func NewCompanyEntity(userID, companyName, companyAddress string, totalStaff int) CompanyEntity {
	now := time.Now()
	return CompanyEntity{
		ID:             uuid.New().String(),
		UserID:         userID,
		CompanyName:    companyName,
		CompanyAddress: companyAddress,
		TotalStaff:     totalStaff,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}
