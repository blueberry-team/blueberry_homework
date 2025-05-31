package user_handler

import (
	"blueberry_homework/dto/request"
	"blueberry_homework/dto/response"
	"blueberry_homework/internal/domain/enum"
	"encoding/json"
	"net/http"
	"strings"
)

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req request.SignUpRequest

	// null check validation
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Email == "" || req.Password == "" || req.Name == "" || req.Role == "" {
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

	// 이메일 유효성 검사
	if !isValidEmail(req.Email) {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   "Invalid email format",
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	// 비밀번호 유효성 검사
	if !isValidPassword(req.Password) {
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   "Invalid password format",
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	// Role 유효성 검사 (boss 또는 worker)
	if !enum.IsUserRoleValid(req.Role) {
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
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response.SuccessResponse{
		Message: "signup successful",
	}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
