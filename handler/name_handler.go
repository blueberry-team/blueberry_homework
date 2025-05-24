package handler

import (
	"net/http"
	"strconv"

	"blueberry_homework_go_gin/usecase"
	"github.com/gin-gonic/gin"
)

// NameHandler 이름 관련 HTTP 요청을 처리하는 핸들러
type NameHandler struct {
	nameUseCase *usecase.NameUseCase
}

// NewNameHandler 새로운 NameHandler 인스턴스를 생성
func NewNameHandler(nameUseCase *usecase.NameUseCase) *NameHandler {
	return &NameHandler{nameUseCase: nameUseCase}
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

	// 이름 길이 검증
	if len(req.Name) < 1 || len(req.Name) > 6 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "name must be between 1 and 6 characters"})
		return
	}

	// 이름 생성
	if err := h.nameUseCase.CreateName(req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": err.Error()})
		return
	}

	// POST 요청 응답: message만 포함
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// GetNames 모든 이름을 조회하는 핸들러
func (h *NameHandler) GetNames(c *gin.Context) {
	names, err := h.nameUseCase.GetNames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error", "error": "failed to get names"})
		return
	}

	// GET 요청 응답: message와 data 포함
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": names})
}

// ChangeName 이름을 변경하는 핸들러
func (h *NameHandler) ChangeName(c *gin.Context) {
	var req ChangeNameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "invalid request format"})
		return
	}

	// 이름 길이 검증
	if len(req.Name) < 1 || len(req.Name) > 6 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "name must be between 1 and 6 characters"})
		return
	}

	// 이름 변경
	if err := h.nameUseCase.ChangeName(req.ID, req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": err.Error()})
		return
	}

	// PUT 요청 응답: message만 포함
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// DeleteByIndex 인덱스로 이름을 삭제하는 핸들러
func (h *NameHandler) DeleteByIndex(c *gin.Context) {
	indexStr := c.Query("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "invalid index"})
		return
	}

	// 이름 삭제
	err = h.nameUseCase.DeleteByIndex(index)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": err.Error()})
		return
	}

	// DELETE 요청 응답: message만 포함
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

// DeleteByName 이름으로 항목을 삭제하는 핸들러
func (h *NameHandler) DeleteByName(c *gin.Context) {
	var req DeleteNameRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "invalid request format"})
		return
	}

	// 이름들 삭제
	err := h.nameUseCase.DeleteByName(req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": err.Error()})
		return
	}

	// DELETE 요청 응답: message만 포함
	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
