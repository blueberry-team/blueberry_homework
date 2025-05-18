package entities

import (
	"time"

	"github.com/gocql/gocql"
)

type CompanyEntity struct {
	Id gocql.UUID `json:"id"`
	Name string `json:"name"`
	CompanyName string `json:"companyName"`
	CreatedAt time.Time `json:"createdAt"`
}
