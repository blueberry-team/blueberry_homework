package repository

import (
	"blueberry_homework/internal/domain/entities"
	"blueberry_homework/internal/domain/repo_interface"
	"fmt"
	"sync"
	"time"

	"github.com/gocql/gocql"
)

// companyRepo CompanyRepository 인터페이스의 구현체입니다.
type companyRepo struct {
	// 저장소
	session *gocql.Session
	// Mutex 추가
	mu sync.Mutex

}

// NewCompanyRepository 새로운 CompanyRepository 인스턴스를 생성합니다.
func NewCompanyRepository(session *gocql.Session) repo_interface.CompanyRepository {
	return &companyRepo{
		session: session,
	}
}

// Company entity 를 저장소에 추가하는 함수
func (r *companyRepo) CreateCompany(entity entities.CompanyEntity) error {
	// Mutex로 락을 걸어 동시 접근을 방지합니다.
	r.mu.Lock()
	defer r.mu.Unlock()

	// 컴패니 중복 확인
	var existingCompany string
	err := r.session.Query(`
		SELECT company_name FROM companies WHERE company_name = ? LIMIT 1
	`, entity.CompanyName).Scan(&existingCompany)
	if err == nil {
		return fmt.Errorf("company already exists: %s", existingCompany)
	}
	if err != gocql.ErrNotFound {
		return err
	}

	// INSERT
	return r.session.Query(`
		INSERT INTO companies (id, name, company_name, created_at)
		VALUES (?, ?, ?, ?)
	`, entity.Id, entity.Name, entity.CompanyName, entity.CreatedAt).Exec()
}

// GetCompanies는 저장된 모든 company 정보를 반환합니다.
func (r *companyRepo) GetCompanies() ([]entities.CompanyEntity, error) {
	iter := r.session.Query(`
		SELECT id, name, company_name, created_at FROM companies
	`).Iter()

	var results []entities.CompanyEntity
	var id gocql.UUID
	var name, companyName string
	var createdAt time.Time

	for iter.Scan(&id, &name, &companyName, &createdAt) {
		results = append(results, entities.CompanyEntity{
			Id:          id,
			Name:        name,
			CompanyName: companyName,
			CreatedAt:   createdAt,
		})
	}

	if err := iter.Close(); err != nil {
		return nil, fmt.Errorf("GetNames() query failed: %v", err)
	}

	return results, nil
}
