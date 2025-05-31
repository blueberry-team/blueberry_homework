package jwt

import (
	"fmt"
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
		"exp":   jwt.NewNumericDate(now.Add(time.Minute * 5)),
		"iat":   jwt.NewNumericDate(now),
	}

	if iat, ok := claims["iat"].(*jwt.NumericDate); ok {
		fmt.Println("발급 시간:", iat.Time.Format("2006-01-02 15:04:05"))
	}
	if exp, ok := claims["exp"].(*jwt.NumericDate); ok {
		fmt.Println("만료 시간:", exp.Time.Format("2006-01-02 15:04:05"))
	}

	tokenString := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	secretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	// 비밀키로 토큰 서명
	tokenSealed, err := tokenString.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenSealed, nil
}
