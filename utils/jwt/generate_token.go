package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken 함수는 사용자 정보를 기반으로 JWT를 생성합니다.
func GenerateToken(userId, email, name string) (string, error) {
	now := time.Now()

	claims := jwt.MapClaims{
		"sub":   userId,
		"email": email,
		"name":  name,
		"exp":   jwt.NewNumericDate(now.Add(time.Hour * 5)),
		"iat":   jwt.NewNumericDate(now),
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte(os.Getenv("JWT_SECRET"))
	// 비밀키로 토큰 서명
	tokenSealed, err := tokenString.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenSealed, nil
}
