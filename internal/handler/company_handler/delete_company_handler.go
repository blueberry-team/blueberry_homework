package company_handler

import (
	"blueberry_homework/dto/response"
	"encoding/json"
	"net/http"
)

func (h *CompanyHandler) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

	err = h.CompanyUsecase.DeleteCompany(userId)
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
		Message: "success",
	}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
