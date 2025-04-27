package handler

import (
	"blueberry_homework/internal/domain/usecase"
	"blueberry_homework/internal/dto"
	"blueberry_homework/internal/response"
	"encoding/json"
	"net/http"
)

type CompanyHandler struct {
	createUsecase *usecase.CreateCompanyUsecase
	companyUsecase *usecase.CompanyUsecase
}

func NewCompanyHandler (cu *usecase.CreateCompanyUsecase, u *usecase.CompanyUsecase) *CompanyHandler {
	return &CompanyHandler{createUsecase: cu, companyUsecase: u}
}

func (h *CompanyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	var req req.CreateCompanyRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res.ErrorResponse{
			Message: "error",
			Error:   "Invalid request format",
		})
		return
	}

	err = h.createUsecase.CreateCompany(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res.ErrorResponse{
			Message: "error",
			Error:   err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res.SuccessResponse{
		Message: "success",
	})
}

func (h *CompanyHandler) GetCompanies(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res.GetCompaniesResponse{
		Message: "success",
		Data:    h.companyUsecase.GetCompanies(),
	})
}
