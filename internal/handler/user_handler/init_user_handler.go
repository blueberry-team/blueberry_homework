// 파일 이동 작업이므로 실제 코드 변경은 없고, 파일만 이동합니다.

package user_handler

import (
	"blueberry_homework/internal/domain/usecase/user_usecase"
	"blueberry_homework/middleware"
	"errors"
	"net/http"
	"regexp"

	"github.com/golang-jwt/jwt/v5"
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

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// 비밀번호 강도 검증 (최소 4자, 숫자, 특수문자 포함)
func isValidPassword(password string) bool {
	return len(password) >= 4
	// regexp.MustCompile(`[a-z]`).MatchString(password) &&
	// regexp.MustCompile(`[A-Z]`).MatchString(password) &&
	// regexp.MustCompile(`[0-9]`).MatchString(password) &&
	// regexp.MustCompile(`[!@#$%^&*]`).MatchString(password)
}

// 미들웨어에서 저장한 context 에서 userId 추출하는 함수
func getUserIdFromContext(r *http.Request) (string, error) {
    claimsValue := r.Context().Value(middleware.ClaimsContextKey)
    if claimsValue == nil {
        return "", errors.New("JWT claims not found in context")
    }
    claims, ok := claimsValue.(jwt.MapClaims)
    if !ok {
        return "", errors.New("failed to parse JWT claims from context")
    }
    userId, ok := claims["sub"].(string)
    if !ok {
        return "", errors.New("invalid user ID in token")
    }
    return userId, nil
}
