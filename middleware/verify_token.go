package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"blueberry_homework/dto/response"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const ClaimsContextKey = contextKey("jwtClaims")

// JWT 토큰 검증 미들웨어
// chi router use 함수가 http.handler() 를 리턴해야함으로 아래와 같은 양식으로 작성됨
func VerifyToken(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// 1. Authorization 헤더에서 토큰 추출
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			if err := json.NewEncoder(w).Encode(response.ErrorResponse{
				Message: "error",
				Error:   "Authorization header is missing",
			}); err != nil {
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			}
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			w.WriteHeader(http.StatusUnauthorized)
			if err := json.NewEncoder(w).Encode(response.ErrorResponse{
				Message: "error",
				Error:   "Invalid Authorization header format. Expected 'Bearer <token>'",
			}); err != nil {
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			}
			return
		}

		// 2. 토큰 파싱 및 서명 검증
		secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
		if len(secretKey) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			if err := json.NewEncoder(w).Encode(response.ErrorResponse{
				Message: "error",
				Error:   "JWT_SECRET_KEY environment variable not set",
			}); err != nil {
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			}
			return
		}

		// 서명 인코딩 알고리즘 확인
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if alg, ok := token.Header["alg"].(string); !ok || alg != jwt.SigningMethodHS512.Alg() {
				return nil, fmt.Errorf("unexpected signing algorithm: %v, expected %v", token.Header["alg"], jwt.SigningMethodHS512.Alg())
			}
			return secretKey, nil
		})
		if err != nil {
			if token != nil && !token.Valid {
				w.WriteHeader(http.StatusUnauthorized)
				if err := json.NewEncoder(w).Encode(response.ErrorResponse{
					Message: "error",
					Error:   "token is invalid",
				}); err != nil {
					http.Error(w, "Failed to encode response", http.StatusInternalServerError)
				}
				return
			}

			w.WriteHeader(http.StatusUnauthorized)
			if err := json.NewEncoder(w).Encode(response.ErrorResponse{
				Message: "error",
				Error:   "token parsing/validation failed: " + err.Error(),
			}); err != nil {
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			}
			return
		}

		// 3. 만료시간(exp) 체크
		// generate_token.go 에서 토큰 생성 시 사용한 커스텀 타입 MapClaims 로 타입 단언
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			if err := json.NewEncoder(w).Encode(response.ErrorResponse{
				Message: "error",
				Error:   "failed to parse token claims as MapClaims",
			}); err != nil {
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			}
			return
		}
		expClaim, exists := claims["exp"] // Type: interface{}
		if !exists {
			w.WriteHeader(http.StatusUnauthorized)
			if err := json.NewEncoder(w).Encode(response.ErrorResponse{
				Message: "error",
				Error:   "'exp' claim not found in token",
			}); err != nil {
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			}
			return
		}
		expFloat, ok := expClaim.(float64) // type: float64
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			if err := json.NewEncoder(w).Encode(response.ErrorResponse{
				Message: "error",
				Error:   "invalid 'exp' claim type",
			}); err != nil {
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			}
			return
		}
		expTime := time.Unix(int64(expFloat), 0) // unix는 int64를 필요로 함
		now := time.Now()
		durationUntilExpiry := expTime.Sub(now)
		if durationUntilExpiry <= 0 {
			w.WriteHeader(http.StatusUnauthorized)
			if err := json.NewEncoder(w).Encode(response.ErrorResponse{
				Message: "error",
				Error:   "token is expired",
			}); err != nil {
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			}
			return
		}
		if durationUntilExpiry < time.Hour {
			w.Header().Set("X-Token-Refresh-Required", "true")
		}

		// 4. claims를 context에 저장
		ctx := context.WithValue(r.Context(), ClaimsContextKey, claims)
		// 5. 다음 handler로 전달
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
