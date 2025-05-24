package db

import (
	"fmt"
	"time"

	"blueberry_homework/config"

	"github.com/gocql/gocql"
)

func InitScylla(cfg *config.Config) (*gocql.Session, error) {
	cluster := gocql.NewCluster(cfg.ScyllaHosts...)

	// Consistency 설정
	cluster.Consistency = gocql.Quorum
	// 연결 타임아웃 설정
	cluster.Timeout = 30 * time.Second
	// 쿼리 타임아웃 설정
	cluster.ConnectTimeout = 10 * time.Second

	// 지정 키스페이스 생성 (블루베리 키스페이스 생성)
	if err := createKeyspace(cluster); err != nil {
		return nil, err
	}

	cluster.Keyspace = cfg.ScyllaKeyspace

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("scylla 연결 실패 (blueberry): %v", err)
	}

	// session.Query(`DROP TABLE IF EXISTS users;`).Exec()
	// session.Query(`DROP TABLE IF EXISTS companies;`).Exec()

	// users 테이블 초기화
	if err := initUserTable(session); err != nil {
		session.Close()
		return nil, err
	}

	// companies 테이블 초기화
	if err := initCompanyTable(session); err != nil {
		session.Close()
		return nil, err
	}

	fmt.Println("✅ Scylla 초기화 완료!")
	return session, nil
}

// createKeyspace blueberry 키스페이스 생성
func createKeyspace(cluster *gocql.ClusterConfig) error {
	// keyspace 생성용 세션
	session, err := cluster.CreateSession()
	if err != nil {
		return fmt.Errorf("scylla 연결 실패 (초기): %v", err)
	}
	defer session.Close()

	// blueberry 키스페이스 생성 (존재 시 무시)
	err = session.Query(`
		CREATE KEYSPACE IF NOT EXISTS blueberry
		WITH replication = {
			'class': 'SimpleStrategy',
			'replication_factor': 1
		};
	`).Exec()
	if err != nil {
		return fmt.Errorf("키스페이스 생성 실패: %v", err)
	}

	return nil
}

// initUserTable users 테이블 및 인덱스 생성
func initUserTable(session *gocql.Session) error {
	// users 테이블 생성
	err := session.Query(`
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY,
			email TEXT,
			password TEXT,
			name TEXT,
			role TEXT,
			created_at TIMESTAMP,
			updated_at TIMESTAMP
		);
	`).Exec()
	if err != nil {
		return fmt.Errorf("users 테이블 생성 실패: %v", err)
	}

	// email 컬럼 인덱스 생성 (FILTER 용도)
	err = session.Query(`
		CREATE INDEX IF NOT EXISTS ON users (email);
	`).Exec()
	if err != nil {
		return fmt.Errorf("email 인덱스 생성 실패: %v", err)
	}

	// role 컬럼 인덱스 생성 (FILTER 용도)
	err = session.Query(`
		CREATE INDEX IF NOT EXISTS ON users (role);
	`).Exec()
	if err != nil {
		return fmt.Errorf("role 인덱스 생성 실패: %v", err)
	}

	return nil
}

// initCompanyTable companies 테이블 및 인덱스 생성
func initCompanyTable(session *gocql.Session) error {
	// companies 테이블 생성
	err := session.Query(`
		CREATE TABLE IF NOT EXISTS companies (
			id UUID PRIMARY KEY,
			user_id UUID,
			company_name TEXT,
			company_address TEXT,
			total_staff INT,
			created_at TIMESTAMP,
			updated_at TIMESTAMP
		);
	`).Exec()
	if err != nil {
		return fmt.Errorf("companies 테이블 생성 실패: %v", err)
	}

	// user_id 컬럼 인덱스 생성 (FILTER 용도)
	err = session.Query(`
		CREATE INDEX IF NOT EXISTS ON companies (user_id);
	`).Exec()
	if err != nil {
		return fmt.Errorf("user_id 인덱스 생성 실패: %v", err)
	}

	return nil
}
