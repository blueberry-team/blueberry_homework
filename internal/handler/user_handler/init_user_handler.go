package user_handler

import (
	"blueberry_homework/internal/domain/usecase/user_usecase"
	"regexp"
)

// UserHandler는 사용자 관련 HTTP 요청을 처리하는 핸들러입니다.
type UserHandler struct {
	usecase *user_usecase.UserUsecase
}

// NewUserHandler는 새로운 UserHandler 인스턴스를 생성합니다.
// usecase.UserUsecase를 의존성으로 주입받습니다.
func NewUserHandler(u *user_usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

var (
	lowercaseRegex = regexp.MustCompile(`[a-z]`)
	uppercaseRegex = regexp.MustCompile(`[A-Z]`)
	digitRegex     = regexp.MustCompile(`[0-9]`)
	specialRegex   = regexp.MustCompile(`[!@#$%^&*]`)
)

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// 비밀번호 강도 검증 (최소 8자, 소문자, 대문자, 숫자, 특수문자 포함)
func isValidPassword(password string) bool {
	return len(password) >= 8 &&
		lowercaseRegex.MatchString(password) &&
		uppercaseRegex.MatchString(password) &&
		digitRegex.MatchString(password) &&
		specialRegex.MatchString(password)
}
