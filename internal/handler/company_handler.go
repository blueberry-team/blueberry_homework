package handler

import (
	"blueberry_homework/dto/request"
	"blueberry_homework/dto/response"
	"blueberry_homework/internal/domain/usecase/company_usecase"
	"encoding/json"
	"net/http"
)

// CompanyHandler는 회사 관련 HTTP 요청을 처리하는 핸들러입니다.
type CompanyHandler struct {
	CompanyUsecase *company_usecase.CompanyUsecase
}

// NewCompanyHandler는 새로운 CompanyHandler 인스턴스를 생성합니다.
// createUsecase와 companyUsecase를 의존성으로 주입받습니다.
func NewCompanyHandler(u *company_usecase.CompanyUsecase) *CompanyHandler {
	return &CompanyHandler{CompanyUsecase: u}
}

// CreateCompany는 새로운 회사를 생성합니다.
func (h *CompanyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var req request.CreateCompanyRequest

	// null check validation
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.CompanyName == "" || req.UserID == "" || req.CompanyAddress == "" {
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

	err = h.CompanyUsecase.CreateCompany(req)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response.SuccessResponse{
		Message: "company created successfully",
	}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// GetUserCompany는 특정 유저의 회사 정보를 조회합니다.
func (h *CompanyHandler) GetUserCompany(w http.ResponseWriter, r *http.Request) {
	var req request.GetCompanyRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.UserId == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   "userId is required",
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	company, err := h.CompanyUsecase.GetUserCompany(req.UserId)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(company); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// ChangeCompany는 회사 정보를 수정합니다.
func (h *CompanyHandler) ChangeCompany(w http.ResponseWriter, r *http.Request) {
	var req request.ChangeCompanyRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.UserId == "" || req.CompanyName == "" || req.CompanyAddress == "" {
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

	err = h.CompanyUsecase.ChangeCompany(req)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response.SuccessResponse{
		Message: "company changed successfully",
	}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// DeleteCompany는 회사를 삭제합니다.
func (h *CompanyHandler) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	var req request.DeleteCompanyRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.UserId == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   "userId is required",
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
		return
	}

	err = h.CompanyUsecase.DeleteCompany(req.UserId)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response.SuccessResponse{
		Message: "success",
	}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
