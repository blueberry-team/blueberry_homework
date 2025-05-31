package user_handler

import (
	"blueberry_homework/dto/response"
	"blueberry_homework/utils/ctxutil"
	"encoding/json"
	"net/http"
)

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get user id from context
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

	// get user data
	user, err := h.usecase.GetUser(userId)
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

	// return user data
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response.GetUserResponse{
		Message: "success",
		Data:    user,
	}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
