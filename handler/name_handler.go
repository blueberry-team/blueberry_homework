package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"blueberry_homework_go_gin/repository"
)

type NameHandler struct {//구조를 만들어놓음
	repo *repository.NameRepository
}

func NewNameHandler(repo *repository.NameRepository) *NameHandler {
	return &NameHandler{repo: repo}
}

type CreateNameRequest struct {
	Name string `json:"name" binding:"required"`
}

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

	h.repo.CreateName(req.Name)
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": h.repo.GetName()})
}

func (h *NameHandler) GetName(c *gin.Context) {
	names := h.repo.GetName()
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": names})
}

func (h *NameHandler) DeleteName(c *gin.Context) {
	indexStr := c.Query("index")
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "invalid index"})
		return
	}

	success := h.repo.DeleteName(index)
	if !success {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "index out of range"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": h.repo.GetName()})
}
