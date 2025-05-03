package handler

import (
	"blueberry_homework/internal/domain/usecase"
	"blueberry_homework/internal/request"
	"blueberry_homework/internal/response"
	"encoding/json"
	"net/http"
)

// CompanyHandler는 회사 관련 HTTP 요청을 처리하는 핸들러입니다.
type CompanyHandler struct {
	createUsecase  *usecase.CreateCompanyUsecase
	companyUsecase *usecase.CompanyUsecase
}

// NewCompanyHandler는 새로운 CompanyHandler 인스턴스를 생성합니다.
// createUsecase와 companyUsecase를 의존성으로 주입받습니다.
func NewCompanyHandler(cu *usecase.CreateCompanyUsecase, u *usecase.CompanyUsecase) *CompanyHandler {
	return &CompanyHandler{createUsecase: cu, companyUsecase: u}
}

// CreateCompany는 새로운 회사를 생성합니다.
func (h *CompanyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var req request.CreateCompanyRequest

	// null check validation
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.CompanyName == "" || req.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   "Invalid request format",
		})
		return
	}

	// 중복 에러 봔환 확인
	err = h.createUsecase.CreateCompany(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response.SuccessResponse{
		Message: "success",
	})
}

// GetCompanies는 모든 회사 목록을 조회하는 HTTP 엔드포인트를 처리합니다.
// 성공 시 200 OK 상태 코드와 함께 회사 목록을 반환합니다.
func (h *CompanyHandler) GetCompanies(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.GetCompaniesResponse{
		Message: "success",
		Data:    h.companyUsecase.GetCompanies(),
	})
}
