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

type GetUserRequest struct {
	Id string `json:"id"`
}

type ChangeUserRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Role enum.UserRole `json:"role"`
}
