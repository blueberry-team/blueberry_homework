package request

import "blueberry_homework/internal/domain/enum"

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Role     enum.UserRole `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangeUserRequest struct {
	Name string `json:"name"`
	Role enum.UserRole `json:"role"`
}
