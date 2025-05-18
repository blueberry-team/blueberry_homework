package entities

import (
	"time"

	"github.com/gocql/gocql"
)

type NameEntity struct {
	Id gocql.UUID `json:"id"`
	Name string `json:"name"`
    CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
