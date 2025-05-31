package auth_handler

import (
	"blueberry_homework/dto/response"
	"blueberry_homework/internal/domain/usecase/auth_usecase"
	"blueberry_homework/utils/ctxutil"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	AuthUsecase *auth_usecase.AuthUsecase
}

func NewAuthHandler(u *auth_usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{AuthUsecase: u}
}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

	token, err := h.AuthUsecase.RefreshToken(userId)
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

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response.GetTokenResponse{
		Message: "token refreshed successfully",
		Data:    token,
	}); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
