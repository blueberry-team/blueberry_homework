package usecase

import (
	"errors"

	"blueberry_homework_go_gin/entity"
	"blueberry_homework_go_gin/repository"
)

// CompanyUseCase 회사 관련 비즈니스 로직을 담당하는 유스케이스
type CompanyUseCase struct {
	companyRepo *repository.CompanyRepository
	nameRepo    *repository.NameRepository
}

// NewCompanyUseCase 새로운 CompanyUseCase 인스턴스를 생성
func NewCompanyUseCase(companyRepo *repository.CompanyRepository, nameRepo *repository.NameRepository) *CompanyUseCase {
	return &CompanyUseCase{
		companyRepo: companyRepo,
		nameRepo:    nameRepo,
	}
}

// CreateCompany 새로운 회사를 생성하고 저장
func (u *CompanyUseCase) CreateCompany(userName, companyName string) error {
	// 1. 사용자 존재 여부 확인
	user, err := u.nameRepo.FindByName(userName)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	// 2. 이미 회사를 가지고 있는지 확인
	existingCompany, err := u.companyRepo.FindCompanyByUserName(userName)
	if err != nil {
		return err
	}
	if existingCompany != nil {
		return errors.New("user already has a company")
	}

	// 3. 새 회사 생성
	company := entity.NewCompanyEntity(userName, companyName)
	return u.companyRepo.CreateCompany(company)
}

// GetAllCompanies 모든 회사 목록을 조회
func (u *CompanyUseCase) GetAllCompanies() ([]entity.CompanyEntity, error) {
	return u.companyRepo.GetAllCompanies()
}

// FindCompanyByUserName 사용자 이름으로 회사 찾기 (내부 로직용)
func (u *CompanyUseCase) FindCompanyByUserName(userName string) (*entity.CompanyEntity, error) {
	return u.companyRepo.FindCompanyByUserName(userName)
}
