package user_handler

import (
	"blueberry_homework/dto/request"
	"blueberry_homework/dto/response"
	"blueberry_homework/internal/domain/enum"
	"encoding/json"
	"net/http"
	"strings"
)

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
	if !enum.IsUserRoleValid(req.Role) {
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
