package ctxutil

import (
	"blueberry_homework/middleware"
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

// 미들웨어에서 저장한 context 에서 userId 추출하는 함수
func GetUserIdFromContext(r *http.Request) (string, error) {
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
