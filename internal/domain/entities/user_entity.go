package entities

import (
	"time"

	"github.com/gocql/gocql"
)

type UserEntity struct {
	Id        gocql.UUID `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Name      string     `json:"name"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}
