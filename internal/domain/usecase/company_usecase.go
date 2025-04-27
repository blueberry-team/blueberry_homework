package usecase

import (
	"blueberry_homework/internal/domain/entities"
	"blueberry_homework/internal/domain/repo_interface"
)

type CompanyUsecase struct {
	repo repointerface.CompanyRepository
}

func NewCompanyUsecase(r repointerface.CompanyRepository) *CompanyUsecase {
	return &CompanyUsecase{
		repo: r,
	}
}

func (u *CompanyUsecase) GetCompanies() []entities.CompanyEntity {
	return u.repo.GetCompanies()
}
