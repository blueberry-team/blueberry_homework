package handler

import (
	"net/http"

	"blueberry_homework_go_gin/usecase"
	"github.com/gin-gonic/gin"
)

// AuthHandler 인증 관련 HTTP 요청을 처리하는 핸들러
type AuthHandler struct {
	authUseCase *usecase.AuthUseCase
}

// NewAuthHandler 새로운 AuthHandler 인스턴스를 생성
func NewAuthHandler(authUseCase *usecase.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUseCase: authUseCase}
}

// SignUpRequest 회원가입 요청 구조체
type SignUpRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

// LogInRequest 로그인 요청 구조체
type LogInRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ChangeUserRequest 사용자 정보 수정 요청 구조체
type ChangeUserRequest struct {
	UserID string `json:"userId" binding:"required"`
	Name   string `json:"name,omitempty"`
	Role   string `json:"role,omitempty"`
}

// GetUserRequest 사용자 정보 조회 요청 구조체
type GetUserRequest struct {
	UserID string `json:"userId" binding:"required"`
}

// SignUp 회원가입 핸들러
func (h *AuthHandler) SignUp(c *gin.Context) {
	var req SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "invalid request format"})
		return
	}

	// 회원가입 처리
	user, err := h.authUseCase.SignUp(req.Email, req.Password, req.Name, req.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": err.Error()})
		return
	}

	// POST 요청 응답: message와 data 포함 (회원가입 성공 시 사용자 정보 반환)
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

// LogIn 로그인 핸들러
func (h *AuthHandler) LogIn(c *gin.Context) {
	var req LogInRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "invalid request format"})
		return
	}

	// 로그인 처리
	user, err := h.authUseCase.LogIn(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "error", "error": err.Error()})
		return
	}

	// POST 요청 응답: message와 data 포함 (로그인 성공 시 사용자 정보 반환)
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

// GetUser 사용자 정보 조회 핸들러
func (h *AuthHandler) GetUser(c *gin.Context) {
	var req GetUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "invalid request format"})
		return
	}

	// 사용자 정보 조회
	user, err := h.authUseCase.GetUser(req.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "error", "error": err.Error()})
		return
	}

	// GET 요청 응답: message와 data 포함
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}

// ChangeUser 사용자 정보 수정 핸들러
func (h *AuthHandler) ChangeUser(c *gin.Context) {
	var req ChangeUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": "invalid request format"})
		return
	}

	// 사용자 정보 수정
	user, err := h.authUseCase.ChangeUser(req.UserID, req.Name, req.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error", "error": err.Error()})
		return
	}

	// PUT 요청 응답: message와 data 포함 (수정된 사용자 정보 반환)
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
}
