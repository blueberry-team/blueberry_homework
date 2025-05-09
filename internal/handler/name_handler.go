package handler

import (
	"blueberry_homework/internal/domain/usecase"
	"blueberry_homework/internal/request"
	"blueberry_homework/internal/response"

	"encoding/json"
	"net/http"
	"strings"
)

// NameHandler는 이름 관련 HTTP 요청을 처리하는 핸들러입니다.
type NameHandler struct {
	usecase *usecase.NameUsecase
}

// NewNameHandler는 새로운 NameHandler 인스턴스를 생성합니다.
// repository.NameRepository를 의존성으로 주입받습니다.
func NewNameHandler(u *usecase.NameUsecase) *NameHandler {
	return &NameHandler{usecase: u}
}

// CreateName은 새로운 이름을 생성하는 HTTP 엔드포인트 핸들러입니다.
// POST 요청의 body에서 JSON 형식의 이름을 받아 저장합니다.
func (h *NameHandler) CreateName(w http.ResponseWriter, r *http.Request) {
	var req request.NameRequest

	// null check validation
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Name == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error: "Invalid request format",
        }); err != nil {
            http.Error(w, "Failed to encode response", http.StatusInternalServerError)
            return
        }
		return
	}

	// 글자수 제한 validation
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

	// 중복 에러 반환 확인
	err = h.usecase.CreateName(name)
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
	if err:= json.NewEncoder(w).Encode(response.SuccessResponse{
		Message: "success",
	}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// GetNames는 저장된 모든 이름을 조회하는 HTTP 엔드포인트 핸들러입니다.
func (h *NameHandler) GetNames(w http.ResponseWriter, r *http.Request) {
	names, err := h.usecase.GetNames()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "Error",
			Error: err.Error(),
		}); err != nil {
            http.Error(w, "Failed to encode response", http.StatusInternalServerError)
            return
        }
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response.GetNamesResponse{
		Message: "success",
		Data:    names,
	}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *NameHandler) DeleteByName(w http.ResponseWriter, r *http.Request) {
	var req request.DeleteByNameRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || strings.TrimSpace(req.Name) == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   "Invalid request format",
		}); err != nil {
            http.Error(w, "Failed to encode response", http.StatusInternalServerError)
            return
        }
		return
	}

	err = h.usecase.DeleteByName(req.Name)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error: err.Error(),
		}); err != nil {
            http.Error(w, "Failed to encode response", http.StatusInternalServerError)
            return
        }
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response.SuccessResponse{
		Message: "success",
	}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// ChangeName 추가
func (h *NameHandler) ChangeName(w http.ResponseWriter, r *http.Request) {
	var req request.ChangeNameRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || strings.TrimSpace(req.Id) == "" || strings.TrimSpace(req.Name) == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   "invalid request format",
		}); err != nil {
            http.Error(w, "Failed to encode response", http.StatusInternalServerError)
            return
        }
		return
	}

	err = h.usecase.ChangeName(req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err:= json.NewEncoder(w).Encode(response.ErrorResponse{
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
		Message: "success",
	}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
