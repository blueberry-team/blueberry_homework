package company_handler

import (
	"blueberry_homework/dto/request"
	"blueberry_homework/dto/response"
	"encoding/json"
	"net/http"
)

func (h *CompanyHandler) ChangeCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req request.ChangeCompanyRequest

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

	userId, err := getUserIdFromContext(r)
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

	err = h.CompanyUsecase.ChangeCompany(userId, req)
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

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response.SuccessResponse{
		Message: "company changed successfully",
	}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
