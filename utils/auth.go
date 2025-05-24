package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	// Argon2 매개변수
	saltLength = 16
	keyLength  = 32
	time       = 1
	memory     = 64 * 1024
	threads    = 4
)

// HashPassword 비밀번호를 해싱하여 저장 가능한 형태로 변환
func HashPassword(password string) (string, error) {
	// 솔트 생성
	salt := make([]byte, saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// Argon2로 해싱
	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLength)

	// 매개변수, 솔트, 해시를 결합하여 저장
	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, memory, time, threads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash))

	return encodedHash, nil
}

// VerifyPassword 비밀번호 검증
func VerifyPassword(password, encodedHash string) bool {
	// 해시에서 매개변수 추출
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false
	}

	var version int
	_, err := fmt.Sscanf(parts[2], "v=%d", &version)
	if err != nil {
		return false
	}

	if version != argon2.Version {
		return false
	}

	var m, t, p uint32
	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &m, &t, &p)
	if err != nil {
		return false
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false
	}

	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false
	}

	// 입력된 비밀번호로 해시 생성
	otherHash := argon2.IDKey([]byte(password), salt, t, m, uint8(p), uint32(len(hash)))

	// 타이밍 공격 방지를 위한 안전한 비교
	return subtle.ConstantTimeCompare(hash, otherHash) == 1
}

// ValidateEmail 이메일 형식 검증
func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// ValidatePassword 비밀번호 강도 검증
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}
	if len(password) > 128 {
		return fmt.Errorf("password must be no more than 128 characters long")
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)

	if !hasUpper || !hasLower || !hasNumber {
		return fmt.Errorf("password must contain at least one uppercase letter, one lowercase letter, and one number")
	}

	return nil
}
