package repository

import (
	"blueberry_homework_go_gin/entity"
)

// CompanyRepository 회사 정보를 저장하고 관리하는 저장소
type CompanyRepository struct {
	companies []entity.CompanyEntity
}

// NewCompanyRepository 새로운 CompanyRepository 인스턴스를 생성
func NewCompanyRepository() *CompanyRepository {
	return &CompanyRepository{
		companies: make([]entity.CompanyEntity, 0),
	}
}

// CreateCompany 새 회사를 추가
func (r *CompanyRepository) CreateCompany(company entity.CompanyEntity) {
	r.companies = append(r.companies, company)
}

// GetCompanies 모든 회사 목록을 반환
func (r *CompanyRepository) GetCompanies() []entity.CompanyEntity {
	return r.companies
}

// FindCompanyByName 사용자 이름으로 회사를 찾음
func (r *CompanyRepository) FindCompanyByName(name string) *entity.CompanyEntity {
	for i, company := range r.companies {
		if company.Name == name {
			return &r.companies[i]
		}
	}
	return nil
}

// FindCompanyByID ID로 회사를 찾음
func (r *CompanyRepository) FindCompanyByID(id string) *entity.CompanyEntity {
	for i, company := range r.companies {
		if company.ID == id {
			return &r.companies[i]
		}
	}
	return nil
}
