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
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   "Invalid request format",
		})
		return
	}

	// 글자수 제한 validation
	name := strings.TrimSpace(req.Name)
	if len(name) < 1 || len(name) > 50 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   "name must be between 1 and 50 characters",
		})
		return
	}

	// 중복 에러 반환 확인
	err = h.usecase.CreateName(name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response.SuccessResponse{
		Message: "success",
	})
}

// GetNames는 저장된 모든 이름을 조회하는 HTTP 엔드포인트 핸들러입니다.
func (h *NameHandler) GetNames(w http.ResponseWriter, r *http.Request) {
	names, err := h.usecase.GetNames()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "Error",
			Error: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.GetNamesResponse{
		Message: "success",
		Data:    names,
	})
}

// DeleteByIndex는 인덱스를 받아 해당하는 이름을 삭제하는 핸들러입니다.
// func (h *NameHandler) DeleteByIndex(w http.ResponseWriter, r *http.Request) {
// 	var req req.DeleteByIndexRequest

// 	// index type validation
// 	err := json.NewDecoder(r.Body).Decode(&req)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(res.ErrorResponse{
// 			Message: "error",
// 			Error:   "invalid request index",
// 		})
// 		return
// 	}

// 	// index range validation
// 	currentNames := h.usecase.GetNames()
// 	if req.Index < 0 || req.Index >= len(currentNames) {
// 		w.WriteHeader(http.StatusBadRequest)
// 		json.NewEncoder(w).Encode(res.ErrorResponse{
// 			Message: "error",
// 			Error:   "invalid index range",
// 		})
// 		return
// 	}

// 	h.usecase.DeleteByIndex(req.Index)

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(res.SuccessResponse{
// 		Message: "success",
// 	})
// }

func (h *NameHandler) DeleteByName(w http.ResponseWriter, r *http.Request) {
	var req request.DeleteByNameRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || strings.TrimSpace(req.Name) == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   "Invalid request format",
		})
		return
	}

	err = h.usecase.DeleteByName(req.Name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.SuccessResponse{
		Message: "success",
	})
}

// ChangeName 추가
func (h *NameHandler) ChangeName(w http.ResponseWriter, r *http.Request) {
	var req request.ChangeNameRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || strings.TrimSpace(req.Id) == "" || strings.TrimSpace(req.Name) == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   "invalid request format",
		})
		return
	}

	err = h.usecase.ChangeName(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response.SuccessResponse{
		Message: "success",
	})
}
