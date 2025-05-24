package usecase

import (
	"errors"

	"blueberry_homework_go_gin/entity"
	"blueberry_homework_go_gin/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// CompanyUseCase 회사 관련 비즈니스 로직을 담당하는 유스케이스
type CompanyUseCase struct {
	companyRepo *repository.CompanyRepository
	authRepo    *repository.AuthRepository
}

// NewCompanyUseCase 새로운 CompanyUseCase 인스턴스를 생성
func NewCompanyUseCase(companyRepo *repository.CompanyRepository, authRepo *repository.AuthRepository) *CompanyUseCase {
	return &CompanyUseCase{
		companyRepo: companyRepo,
		authRepo:    authRepo,
	}
}

// CreateCompany 새로운 회사를 생성하고 저장 (boss만 가능)
func (u *CompanyUseCase) CreateCompany(userID, companyName, companyAddress string, totalStaff int) (*entity.CompanyEntity, error) {
	// 1. 사용자 존재 여부 및 권한 확인
	user, err := u.authRepo.FindUserByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// 2. boss 권한 확인
	if !user.IsBoss() {
		return nil, errors.New("only boss can create company")
	}

	// 3. 이미 회사를 가지고 있는지 확인
	existingCompany, err := u.companyRepo.FindCompanyByUserID(userID)
	if err != nil {
		return nil, err
	}
	if existingCompany != nil {
		return nil, errors.New("user already has a company")
	}

	// 4. 입력 값 검증
	if companyName == "" {
		return nil, errors.New("company name is required")
	}
	if companyAddress == "" {
		return nil, errors.New("company address is required")
	}
	if totalStaff < 0 {
		return nil, errors.New("total staff must be non-negative")
	}

	// 5. 새 회사 생성
	company := entity.NewCompanyEntity(userID, companyName, companyAddress, totalStaff)
	if err := u.companyRepo.CreateCompany(company); err != nil {
		return nil, err
	}

	return &company, nil
}

// GetAllCompanies 모든 회사 목록을 조회
func (u *CompanyUseCase) GetAllCompanies() ([]entity.CompanyEntity, error) {
	return u.companyRepo.GetAllCompanies()
}

// GetCompanyByID ID로 회사 조회
func (u *CompanyUseCase) GetCompanyByID(companyID string) (*entity.CompanyEntity, error) {
	company, err := u.companyRepo.FindCompanyByID(companyID)
	if err != nil {
		return nil, err
	}
	if company == nil {
		return nil, errors.New("company not found")
	}
	return company, nil
}

// GetCompanyByUserID 사용자 ID로 회사 조회
func (u *CompanyUseCase) GetCompanyByUserID(userID string) (*entity.CompanyEntity, error) {
	company, err := u.companyRepo.FindCompanyByUserID(userID)
	if err != nil {
		return nil, err
	}
	if company == nil {
		return nil, errors.New("company not found")
	}
	return company, nil
}

// ChangeCompany 회사 정보 수정 (소유자만 가능)
func (u *CompanyUseCase) ChangeCompany(userID, companyID, companyName, companyAddress string, totalStaff *int) (*entity.CompanyEntity, error) {
	// 1. 회사 존재 확인
	company, err := u.companyRepo.FindCompanyByID(companyID)
	if err != nil {
		return nil, err
	}
	if company == nil {
		return nil, errors.New("company not found")
	}

	// 2. 소유자 권한 확인
	if company.UserID != userID {
		return nil, errors.New("only company owner can modify company information")
	}

	// 3. 업데이트할 필드 준비
	updates := bson.D{}

	if companyName != "" {
		updates = append(updates, bson.E{Key: "companyName", Value: companyName})
	}

	if companyAddress != "" {
		updates = append(updates, bson.E{Key: "companyAddress", Value: companyAddress})
	}

	if totalStaff != nil {
		if *totalStaff < 0 {
			return nil, errors.New("total staff must be non-negative")
		}
		updates = append(updates, bson.E{Key: "totalStaff", Value: *totalStaff})
	}

	// 4. 업데이트 실행
	if err := u.companyRepo.UpdateCompany(companyID, updates); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("company not found")
		}
		return nil, err
	}

	// 5. 업데이트된 회사 정보 반환
	return u.GetCompanyByID(companyID)
}

// DeleteCompany 회사 삭제 (소유자만 가능)
func (u *CompanyUseCase) DeleteCompany(userID, companyID string) error {
	// 1. 회사 존재 확인
	company, err := u.companyRepo.FindCompanyByID(companyID)
	if err != nil {
		return err
	}
	if company == nil {
		return errors.New("company not found")
	}

	// 2. 소유자 권한 확인
	if company.UserID != userID {
		return errors.New("only company owner can delete company")
	}

	// 3. 회사 삭제
	if err := u.companyRepo.DeleteCompany(companyID); err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("company not found")
		}
		return err
	}

	return nil
}

// FindCompaniesByName 회사명으로 회사 검색
func (u *CompanyUseCase) FindCompaniesByName(companyName string) ([]entity.CompanyEntity, error) {
	if companyName == "" {
		return nil, errors.New("company name is required for search")
	}
	return u.companyRepo.FindCompaniesByName(companyName)
}
