package company_handler

import (
	"blueberry_homework/dto/request"
	"blueberry_homework/dto/response"
	"encoding/json"
	"net/http"
)

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
