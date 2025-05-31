package company_handler

import (
	"blueberry_homework/internal/domain/usecase/company_usecase"
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
