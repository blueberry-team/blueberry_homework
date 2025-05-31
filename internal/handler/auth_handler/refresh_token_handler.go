package auth_handler

import (
	"blueberry_homework/dto/response"
	"blueberry_homework/internal/domain/usecase/auth_usecase"
	"blueberry_homework/middleware"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	AuthUsecase *auth_usecase.AuthUsecase
}

func NewAuthHandler(u *auth_usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{AuthUsecase: u}
}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
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

	token, err := h.AuthUsecase.RefreshToken(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
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


// 미들웨어에서 저장한 context 에서 userId 추출하는 함수
func getUserIdFromContext(r *http.Request) (string, error) {
    claimsValue := r.Context().Value(middleware.ClaimsContextKey)
    if claimsValue == nil {
        return "", errors.New("JWT claims not found in context")
    }
    claims, ok := claimsValue.(jwt.MapClaims)
    if !ok {
        return "", errors.New("failed to parse JWT claims from context")
    }
    userId, ok := claims["sub"].(string)
    if !ok {
        return "", errors.New("invalid user ID in token")
    }
    return userId, nil
}
