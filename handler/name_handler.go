package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"blueberry_homework_go_gin/repository"
)

type NameHandler struct {
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.repo.CreateName(req.Name)
	c.JSON(http.StatusOK, gin.H{"message": "Name created successfully"})
}

func (h *NameHandler) GetName(c *gin.Context) {
	names := h.repo.GetName()
	c.JSON(http.StatusOK, gin.H{"names": names})
}
