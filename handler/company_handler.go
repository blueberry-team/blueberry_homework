package handler

import (
	"net/http"

	"blueberry_homework_go_gin/usecase"
	"github.com/gin-gonic/gin"
)

// CompanyHandler 회사 관련 HTTP 요청을 처리하는 핸들러
type CompanyHandler struct {
	companyUseCase *usecase.CompanyUseCase
}

// NewCompanyHandler 새로운 CompanyHandler 인스턴스를 생성
func NewCompanyHandler(companyUseCase *usecase.CompanyUseCase) *CompanyHandler {
	return &CompanyHandler{companyUseCase: companyUseCase}
}

// CreateCompanyRequest 회사 생성 요청 구조체
type CreateCompanyRequest struct {
	Name        string `json:"name" binding:"required"`
	CompanyName string `json:"company_name" binding:"required"`
}

// CreateCompany 새 회사를 생성하는 핸들러
func (h *CompanyHandler) CreateCompany(c *gin.Context) {
	var req CreateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "invalid request format"})
		return
	}

	// 회사 생성
	if err := h.companyUseCase.CreateCompany(req.Name, req.CompanyName); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": err.Error()})
		return
	}

	// POST 요청 응답: message만 포함
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// GetAllCompanies 모든 회사를 조회하는 핸들러
func (h *CompanyHandler) GetAllCompanies(c *gin.Context) {
	companies, err := h.companyUseCase.GetAllCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": "failed to get companies"})
		return
	}

	// GET 요청 응답: message와 data 포함
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": companies})
}
