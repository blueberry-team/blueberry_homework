package handler

import (
	"net/http"

	"blueberry_homework_go_gin/entity"
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
	UserID         string `json:"userId" binding:"required"`
	CompanyName    string `json:"companyName" binding:"required"`
	CompanyAddress string `json:"companyAddress" binding:"required"`
	TotalStaff     int    `json:"totalStaff" binding:"required"`
}

// ChangeCompanyRequest 회사 정보 수정 요청 구조체
type ChangeCompanyRequest struct {
	UserID         string `json:"userId" binding:"required"`
	CompanyID      string `json:"companyId" binding:"required"`
	CompanyName    string `json:"companyName,omitempty"`
	CompanyAddress string `json:"companyAddress,omitempty"`
	TotalStaff     *int   `json:"totalStaff,omitempty"`
}

// DeleteCompanyRequest 회사 삭제 요청 구조체
type DeleteCompanyRequest struct {
	UserID    string `json:"userId" binding:"required"`
	CompanyID string `json:"companyId" binding:"required"`
}

// GetCompanyRequest 회사 조회 요청 구조체
type GetCompanyRequest struct {
	CompanyID string `json:"companyId,omitempty"`
	UserID    string `json:"userId,omitempty"`
}

// CreateCompany 새 회사를 생성하는 핸들러
func (h *CompanyHandler) CreateCompany(c *gin.Context) {
	var req CreateCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "invalid request format"})
		return
	}

	// 회사 생성
	company, err := h.companyUseCase.CreateCompany(req.UserID, req.CompanyName, req.CompanyAddress, req.TotalStaff)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": err.Error()})
		return
	}

	// POST 요청 응답: message와 data 포함 (생성된 회사 정보 반환)
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": company})
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

// GetCompany 특정 회사 조회 핸들러
func (h *CompanyHandler) GetCompany(c *gin.Context) {
	var req GetCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "invalid request format"})
		return
	}

	var company *entity.CompanyEntity
	var err error

	// CompanyID 또는 UserID로 조회
	if req.CompanyID != "" {
		company, err = h.companyUseCase.GetCompanyByID(req.CompanyID)
	} else if req.UserID != "" {
		company, err = h.companyUseCase.GetCompanyByUserID(req.UserID)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "companyId or userId is required"})
		return
	}

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "error", "error": err.Error()})
		return
	}

	// GET 요청 응답: message와 data 포함
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": company})
}

// ChangeCompany 회사 정보를 수정하는 핸들러
func (h *CompanyHandler) ChangeCompany(c *gin.Context) {
	var req ChangeCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "invalid request format"})
		return
	}

	// 회사 정보 수정
	company, err := h.companyUseCase.ChangeCompany(req.UserID, req.CompanyID, req.CompanyName, req.CompanyAddress, req.TotalStaff)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": err.Error()})
		return
	}

	// PUT 요청 응답: message와 data 포함 (수정된 회사 정보 반환)
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": company})
}

// DeleteCompany 회사를 삭제하는 핸들러
func (h *CompanyHandler) DeleteCompany(c *gin.Context) {
	var req DeleteCompanyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "invalid request format"})
		return
	}

	// 회사 삭제
	err := h.companyUseCase.DeleteCompany(req.UserID, req.CompanyID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": err.Error()})
		return
	}

	// DELETE 요청 응답: message만 포함
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// FindCompaniesByName 회사명으로 회사를 검색하는 핸들러
func (h *CompanyHandler) FindCompaniesByName(c *gin.Context) {
	companyName := c.Query("name")
	if companyName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "company name is required"})
		return
	}

	companies, err := h.companyUseCase.FindCompaniesByName(companyName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": err.Error()})
		return
	}

	// GET 요청 응답: message와 data 포함
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": companies})
}
