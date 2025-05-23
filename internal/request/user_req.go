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

// GetUserRequest는 사용자 ID를 요청 바디로 받기 위한 구조체입니다.
type GetUserRequest struct {
	Id string `json:"id"`
}

type ChangeUserRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}
