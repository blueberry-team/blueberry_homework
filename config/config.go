package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Config는 애플리케이션 설정을 저장하는 구조체입니다.
type Config struct {
	ScyllaHosts    []string
	ScyllaKeyspace string
	ServerPort     string
}

// NewConfig는 환경변수에서 설정을 로드하여 새 Config 인스턴스를 반환합니다.
func NewConfig() (*Config, error) {
	// .env 파일 로드
	if err := godotenv.Load(); err != nil {
		// .env 파일이 없는 경우 경고만 출력하고 계속 진행
		fmt.Printf("Warning: .env 파일을 찾을 수 없습니다: %v\n", err)
	}

	// ScyllaDB 호스트 설정
	hostsEnv := getEnvOrDefault("SCYLLA_HOSTS", "localhost")

	// 키스페이스 설정
	keyspace := getEnvOrDefault("SCYLLA_KEYSPACE", "blueberry")

	// 서버 포트 설정
	serverPort := getEnvOrDefault("SERVER_PORT", "8080")

	fmt.Println("✅ Config 설정 완료!")

	return &Config{
		ScyllaHosts:    strings.Split(hostsEnv, ","),
		ScyllaKeyspace: keyspace,
		ServerPort:     serverPort,
	}, nil
}

func getEnvOrDefault(key string, defaultValue string) string {
	env := os.Getenv(key)
	if env == "" {
		return defaultValue
	}
	return env
}
