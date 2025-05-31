package entities

import (
	"time"

	"github.com/gocql/gocql"
)

type CompanyEntity struct {
	Id             gocql.UUID `json:"id"`
	UserId         gocql.UUID `json:"userId"`
	CompanyName    string     `json:"companyName"`
	CompanyAddress string     `json:"companyAddress"`
	TotalStaff     int        `json:"totalStaff"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
}

type ChangeCompanyEntity struct {
	UserId         gocql.UUID `json:"userId"`
	CompanyName    string     `json:"companyName"`
	CompanyAddress string     `json:"companyAddress"`
	TotalStaff     int        `json:"totalStaff"`
	UpdatedAt      time.Time  `json:"updatedAt"`
}
