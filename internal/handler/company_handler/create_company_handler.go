package company_handler

import (
	"blueberry_homework/dto/request"
	"blueberry_homework/dto/response"
	"blueberry_homework/utils/ctxutil"
	"encoding/json"
	"net/http"
)

func (h *CompanyHandler) CreateCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req request.CreateCompanyRequest

	// null check validation
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.CompanyName == "" || req.CompanyAddress == "" {
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

	userId, err := ctxutil.GetUserIdFromContext(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if err := json.NewEncoder(w).Encode(response.ErrorResponse{
			Message: "error",
			Error:   err.Error(),
		}); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		}
		return
	}

	err = h.CompanyUsecase.CreateCompany(userId, req)
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

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response.SuccessResponse{
		Message: "company created successfully",
	}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
