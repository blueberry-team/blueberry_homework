package company_handler

import (
	"blueberry_homework/dto/request"
	"blueberry_homework/dto/response"
	"encoding/json"
	"net/http"
)

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
