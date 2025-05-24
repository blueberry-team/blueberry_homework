package handler

import (
	"net/http"
	"strconv"

	"blueberry_homework_go_gin/usecase"
	"github.com/gin-gonic/gin"
)

// NameHandler 이름 관련 HTTP 요청을 처리하는 핸들러
type NameHandler struct {
	useCase *usecase.NameUseCase
}

// NewNameHandler 새로운 NameHandler 인스턴스를 생성
func NewNameHandler(useCase *usecase.NameUseCase) *NameHandler {
	return &NameHandler{useCase: useCase}
}

// CreateNameRequest 이름 생성 요청 구조체
type CreateNameRequest struct {
	Name string `json:"name" binding:"required"`
}

// ChangeNameRequest 이름 변경 요청 구조체
type ChangeNameRequest struct {
	ID   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

// DeleteNameRequest 이름으로 삭제하는 요청 구조체
type DeleteNameRequest struct {
	Name string `json:"name" binding:"required"`
}

// CreateName 새 이름을 생성하는 핸들러
func (h *NameHandler) CreateName(c *gin.Context) {
	var req CreateNameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "invalid request format"})
		return
	}

	if len(req.Name) < 1 || len(req.Name) > 6 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "name must be between 1 and 6 characters"})
		return
	}

	if err := h.useCase.CreateName(req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// ChangeName 이름을 변경하는 핸들러
func (h *NameHandler) ChangeName(c *gin.Context) {
	var req ChangeNameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "invalid request format"})
		return
	}

	if len(req.Name) < 1 || len(req.Name) > 6 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "name must be between 1 and 6 characters"})
		return
	}

	if err := h.useCase.ChangeName(req.ID, req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// GetNames 모든 이름을 조회하는 핸들러
func (h *NameHandler) GetNames(c *gin.Context) {
	names := h.useCase.GetNames()
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": names})
}

// DeleteByIndex 인덱스로 이름을 삭제하는 핸들러
func (h *NameHandler) DeleteByIndex(c *gin.Context) {
	indexStr := c.Query("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "invalid index"})
		return
	}

	success := h.useCase.DeleteByIndex(index)
	if !success {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "index out of range"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// DeleteByName 이름으로 항목을 삭제하는 핸들러
func (h *NameHandler) DeleteByName(c *gin.Context) {
	var req DeleteNameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "invalid request format"})
		return
	}

	success := h.useCase.DeleteByName(req.Name)
	if !success {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "name not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
