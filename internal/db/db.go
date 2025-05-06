package db

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gocql/gocql"
)

func InitScylla() (*gocql.Session, error) {
	// 환경 변수나 설정 파일에서 읽은 값으로 대체
	hosts := getEnvOrDefault("SCYLLA_HOSTS", "localhost")
	cluster := gocql.NewCluster(strings.Split(hosts, ",")...)
	cluster.Consistency = gocql.Quorum
	// 연결 타임아웃 설정
	cluster.Timeout = 10 * time.Second
	// 쿼리 타임아웃 설정
	cluster.ConnectTimeout = 5 * time.Second

	// keyspace 생성용 세션
	session, err := cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("scylla 연결 실패 (초기): %v", err)
	}

	// 2. blueberry 키스페이스 생성 (존재 시 무시)
	err = session.Query(`
		CREATE KEYSPACE IF NOT EXISTS blueberry
		WITH replication = {
			'class': 'SimpleStrategy',
			'replication_factor': 1
		};
	`).Exec()
	if err != nil {
		session.Close()
		return nil, fmt.Errorf("키스페이스 생성 실패: %v", err)
	}
	session.Close() // 초기 세션 종료

	cluster.Keyspace = "blueberry"
	session, err = cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("scylla 연결 실패 (blueberry): %v", err)
	}

	// 4. names 테이블 생성
	err = session.Query(`
		CREATE TABLE IF NOT EXISTS names (
			id UUID PRIMARY KEY,
			name TEXT,
			created_at TIMESTAMP,
			updated_at TIMESTAMP
		);
	`).Exec()
	if err != nil {
		session.Close()
		return nil, fmt.Errorf("names 테이블 생성 실패: %v", err)
	}

	// 5. name 컬럼 인덱스 생성 (FILTER 용도)
	err = session.Query(`
		CREATE INDEX IF NOT EXISTS ON names (name);
	`).Exec()
	if err != nil {
		session.Close()
		return nil, fmt.Errorf("name 인덱스 생성 실패: %v", err)
	}

	// 6. companies 테이블 생성
	err = session.Query(`
		CREATE TABLE IF NOT EXISTS companies (
			id UUID PRIMARY KEY,
			name TEXT,
			company_name TEXT,
			created_at TIMESTAMP
		);
	`).Exec()
	if err != nil {
		session.Close()
		return nil, fmt.Errorf("companies 테이블 생성 실패: %v", err)
	}

	// 7. company_name 컬럼 인덱스 생성 (FILTER 용도)
	err = session.Query(`
		CREATE INDEX IF NOT EXISTS ON companies (company_name);
	`).Exec()
	if err != nil {
		session.Close()
		return nil, fmt.Errorf("company_name 인덱스 생성 실패: %v", err)
	}

	fmt.Println("✅ Scylla 초기화 완료!")
	return session, nil
}

// getEnvOrDefault 환경 변수를 읽거나 기본값을 반환합니다
func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
