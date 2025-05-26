package entities

import (
	"blueberry_homework/internal/domain/enum"
	"time"

	"github.com/gocql/gocql"
)

type UserEntity struct {
	Id        gocql.UUID `json:"id"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Name      string     `json:"name"`
	Role      enum.UserRole     `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
}
