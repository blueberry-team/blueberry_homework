package repository

import (
	"blueberry_homework/internal/domain/entities"
	"blueberry_homework/internal/domain/repo_interface"
	"fmt"
)

type companyRepo struct {
	companies []entities.CompanyEntity
}

func NewCompanyRepository() repointerface.CompanyRepository {
	return &companyRepo{
		companies: []entities.CompanyEntity{},
	}
}

func (r *companyRepo) CreateCompany(entity entities.CompanyEntity) error {
	for _, company := range r.companies {
		if company.Name == entity.Name {
			return fmt.Errorf("company already exist")
		}
	}
	r.companies = append(r.companies, entity)
	return nil
}

func (r *companyRepo) GetCompanies() []entities.CompanyEntity {
	return r.companies
}
