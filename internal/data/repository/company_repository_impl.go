package repository

import (
	"blueberry_homework/internal/domain/entities"
	"blueberry_homework/internal/domain/repo_interface"
	"fmt"

	"github.com/gocql/gocql"
)

// companyRepo CompanyRepository 인터페이스의 구현체입니다.
type companyRepo struct {
	// 저장소
	companies []entities.CompanyEntity
	session *gocql.Session
}

// NewCompanyRepository 새로운 CompanyRepository 인스턴스를 생성합니다.
func NewCompanyRepository(session *gocql.Session) repo_interface.CompanyRepository {
	return &companyRepo{
		companies: []entities.CompanyEntity{},
		session: session,
	}
}

// Company entity 를 저장소에 추가하는 함수
func (r *companyRepo) CreateCompany(entity entities.CompanyEntity) error {
	// 컴패니 중복 확인
	// 추후 기존 컴패니에 이름만 추가하는 방식으로 갈까 싶어서 일단 추가해둠
	for _, company := range r.companies {
		if company.Name == entity.Name {
			return fmt.Errorf("company already exist")
		}
	}
	r.companies = append(r.companies, entity)
	return nil
}

// GetCompanies는 저장된 모든 company 정보를 반환합니다.
func (r *companyRepo) GetCompanies() []entities.CompanyEntity {
	return r.companies
}
