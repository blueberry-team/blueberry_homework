package usecase

import (
	"errors"
	"blueberry_homework_go_gin/entity"
	"blueberry_homework_go_gin/repository"
)

// CompanyUseCase 회사 관련 비즈니스 로직을 담당하는 유스케이스
type CompanyUseCase struct {
	repo     *repository.CompanyRepository
	nameRepo *repository.NameRepository
}

// NewCompanyUseCase 새로운 CompanyUseCase 인스턴스를 생성
func NewCompanyUseCase(repo *repository.CompanyRepository, nameRepo *repository.NameRepository) *CompanyUseCase {
	return &CompanyUseCase{
		repo:     repo,
		nameRepo: nameRepo,
	}
}

// CreateCompany 새로운 회사를 생성하고 저장
func (u *CompanyUseCase) CreateCompany(name, companyName string) error {
	// 사용자 존재 여부 확인
	user := u.nameRepo.FindByName(name)
	if user == nil {
		return errors.New("user not found")
	}

	// 이미 회사를 가지고 있는지 확인
	existingCompany := u.repo.FindCompanyByName(name)
	if existingCompany != nil {
		return errors.New("user already has a company")
	}

	// 새 회사 생성
	company := entity.NewCompanyEntity(name, companyName)
	u.repo.CreateCompany(company)
	return nil
}

// GetCompanies 모든 회사 목록을 조회
func (u *CompanyUseCase) GetCompanies() []entity.CompanyEntity {
	return u.repo.GetCompanies()
}

// FindCompanyByName 이름으로 회사 찾기
func (u *CompanyUseCase) FindCompanyByName(name string) *entity.CompanyEntity {
	return u.repo.FindCompanyByName(name)
}
