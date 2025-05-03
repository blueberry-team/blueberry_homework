package usecase

import (
	"blueberry_homework/internal/domain/entities"
	"blueberry_homework/internal/domain/repo_interface"
)

type CompanyUsecase struct {
	repo repo_interface.CompanyRepository
}

func NewCompanyUsecase(r repo_interface.CompanyRepository) *CompanyUsecase {
	return &CompanyUsecase{
		repo: r,
	}
}

// 저장된 회사 정보 반환 함수
func (u *CompanyUsecase) GetCompanies() []entities.CompanyEntity {
	return u.repo.GetCompanies()
}
