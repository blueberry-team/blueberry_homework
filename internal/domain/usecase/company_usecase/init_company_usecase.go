package company_usecase

import "blueberry_homework/internal/domain/repo_interface"

type CompanyUsecase struct {
	repo repo_interface.CompanyRepository
}

func (u *CompanyUsecase) NewUserCompanyUsecase(companyRepo repo_interface.CompanyRepository, userRepo repo_interface.UserRepository) any {
	panic("unimplemented")
}

func NewCompanyUsecase(r repo_interface.CompanyRepository) *CompanyUsecase {
	return &CompanyUsecase{
		repo: r,
	}
}

type UserCompanyUsecase struct {
	companyRepo repo_interface.CompanyRepository
	userRepo    repo_interface.UserRepository
}

func NewUserCompanyUsecase(companyRepo repo_interface.CompanyRepository, userRepo repo_interface.UserRepository) *UserCompanyUsecase {
	return &UserCompanyUsecase{
		companyRepo: companyRepo,
		userRepo:    userRepo,
	}
}
