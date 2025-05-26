package response

import (
	"blueberry_homework/internal/domain/enum"
	"time"
)

type UserResponse struct {
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Role      enum.UserRole    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserResponse struct {
	Message string       `json:"message"`
	Data    UserResponse `json:"data"`
}
