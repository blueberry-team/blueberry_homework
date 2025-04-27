package entities

import "time"

type CompanyEntity struct {
	Id string `json:"id"`
	Name string `json:"name"`
	CompanyName string `json:"companyName"`
	CreatedAt time.Time `json:"createdAt"`
}
