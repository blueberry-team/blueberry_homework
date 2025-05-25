package handler

import (
	"blueberry_homework/dto/request"
	"blueberry_homework/dto/response"
	"blueberry_homework/internal/domain/usecase/user_usecase"

	"encoding/json"
	"net/http"
	"strings"
)

// UserHandler는 사용자 관련 HTTP 요청을 처리하는 핸들러입니다.
type UserHandler struct {
	usecase *user_usecase.UserUsecase
}

// NewUserHandler는 새로운 UserHandler 인스턴스를 생성합니다.
// usecase.UserUsecase를 의존성으로 주입받습니다.
func NewUserHandler(u *user_usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

// SignUp은 새로운 사용자를 등록하는 HTTP 엔드포인트 핸들러입니다.
func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req request.SignUpRequest

	// null check validation
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Email == "" || req.Password == "" || req.Name == "" || req.Role == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   "Invalid request format or missing fields (Email, Password, Name, Role are required)",
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	// Role 유효성 검사 (boss 또는 worker)
	if req.Role != "boss" && req.Role != "worker" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   "Invalid role. Role must be 'boss' or 'worker'",
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	// 글자수 제한 validation (사용자 이름)
	userName := strings.TrimSpace(req.Name)
	if len(userName) < 1 || len(userName) > 50 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   "name must be between 1 and 50 characters",
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	// 유즈케이스 호출
	err = h.usecase.SignUp(req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   err.Error(),
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	// 성공응답
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response.SuccessResponse{
		Message: "signup successful",
	}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// Login은 사용자 로그인을 처리하는 HTTP 엔드포인트 핸들러입니다.
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req request.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Email == "" || req.Password == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   "Invalid request format or missing fields (Email, Password are required)",
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	success, err := h.usecase.Login(req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "invalid password") {
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   err.Error(),
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	if !success {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized) // 로그인 실패 (자격 증명 불일치 등)
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   "Login failed: Invalid email or password",
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	// 로그인 성공 (향후 JWT 토큰 등을 발급할 수 있습니다)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response.SuccessResponse{
		Message: "login successful",
		// Data: map[string]string{"token": "generated_jwt_token_here"}, // 예시: 토큰 반환
	}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// GetUser는 ID로 특정 사용자를 조회하는 HTTP 엔드포인트 핸들러입니다.
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	var req request.GetUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Id == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   "Invalid request format of json or missing User ID",
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	userRes, err := h.usecase.GetUser(req.Id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(err.Error(), "not found") {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "Error",
			Error:   err.Error(),
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response.GetUserResponse{
		Message: "success",
		Data:    userRes,
	}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// ChangeUser는 기존 사용자 정보를 변경하는 HTTP 엔드포인트 핸들러입니다.
func (h *UserHandler) ChangeUser(w http.ResponseWriter, r *http.Request) {
	var req request.ChangeUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	// ChangeUserRequest DTO의 필드 (Id, Name, Role 등)에 따라 유효성 검사 필요
	if err != nil || req.Id == "" || req.Name == "" || req.Role == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   "invalid request format (Id, Name and Role are required)",
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	// Name 필드에 대한 글자 수 제한 validation (필요하다면)
	name := strings.TrimSpace(req.Name)
	if len(name) < 1 || len(name) > 50 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   "name must be between 1 and 50 characters",
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	// Role 유효성 검사 (boss 또는 worker)
	if req.Role != "boss" && req.Role != "worker" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   "Invalid role. Role must be 'boss' or 'worker'",
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	err = h.usecase.ChangeUser(req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		// 에러 종류에 따라 상태 코드 변경 (예: 사용자를 찾지 못하면 http.StatusNotFound)
		if strings.Contains(err.Error(), "not found") {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusBadRequest) // 혹은 다른 적절한 에러 코드
		}
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   err.Error(),
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response.SuccessResponse{
		Message: "user changed successfully",
	}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
