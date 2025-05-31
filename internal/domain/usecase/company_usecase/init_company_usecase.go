package company_usecase

import "blueberry_homework/internal/domain/repo_interface"

type CompanyUsecase struct {
	companyRepo repo_interface.CompanyRepository
	userRepo    repo_interface.UserRepository
}

func NewCompanyUsecase(companyRepo repo_interface.CompanyRepository, userRepo repo_interface.UserRepository) *CompanyUsecase {
	return &CompanyUsecase{
		companyRepo: companyRepo,
		userRepo:    userRepo,
	}
}
