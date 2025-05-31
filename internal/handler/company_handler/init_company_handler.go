// 파일 이동 작업이므로 실제 코드 변경은 없고, 파일만 이동합니다.
package company_handler

import (
	"blueberry_homework/internal/domain/usecase/company_usecase"
	"blueberry_homework/middleware"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

// CompanyHandler는 회사 관련 HTTP 요청을 처리하는 핸들러입니다.
type CompanyHandler struct {
	CompanyUsecase *company_usecase.CompanyUsecase
}

// NewCompanyHandler는 새로운 CompanyHandler 인스턴스를 생성합니다.
// createUsecase와 companyUsecase를 의존성으로 주입받습니다.
func NewCompanyHandler(u *company_usecase.CompanyUsecase) *CompanyHandler {
	return &CompanyHandler{CompanyUsecase: u}
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
