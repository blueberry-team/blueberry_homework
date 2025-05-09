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
	hostsEnv := os.Getenv("SCYLLA_HOSTS")
	if hostsEnv == "" {
		return nil, fmt.Errorf("SCYLLA_HOSTS 환경변수가 설정되지 않았습니다")
	}

	// 키스페이스 설정
	keyspace := os.Getenv("SCYLLA_KEYSPACE")
	if keyspace == "" {
		return nil, fmt.Errorf("SCYLLA_KEYSPACE 환경변수가 설정되지 않았습니다")
	}

	// 서버 포트 설정
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		return nil, fmt.Errorf("SERVER_PORT 환경변수가 설정되지 않았습니다")
	}
	fmt.Println("✅ Config 설정 완료!")

	return &Config{
		ScyllaHosts:    strings.Split(hostsEnv, ","),
		ScyllaKeyspace: keyspace,
		ServerPort:     serverPort,
	}, nil
}
