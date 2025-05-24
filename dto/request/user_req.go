package request

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ChangeUserRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}
