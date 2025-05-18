package usecase

import (
	"blueberry_homework/internal/domain/entities"
	"blueberry_homework/internal/domain/repo_interface"
	req "blueberry_homework/internal/request"
	"fmt"
	"time"

	"github.com/gocql/gocql"
)

// CreateCompanyUsecase는 회사 생성을 담당하는 유스케이스 구조체입니다.
type CreateCompanyUsecase struct {
	nameRepo    repo_interface.NameRepository
	companyRepo repo_interface.CompanyRepository
}

// NewCreateCompanyUsecase는 CreateCompanyUsecase의 새로운 인스턴스를 생성합니다.
// nameRepo: 이름 조회를 위한 레포지토리
// companyRepo: 회사 정보 관리를 위한 레포지토리
func NewCreateCompanyUsecase(n repo_interface.NameRepository, c repo_interface.CompanyRepository) *CreateCompanyUsecase {
	return &CreateCompanyUsecase{
		nameRepo:    n,
		companyRepo: c,
	}
}

// CreateCompany는 새로운 회사를 생성합니다.
// 레포에서 중복확인이 이뤄지기 떄문에 함수의 결과값을 리턴합니다
func (cr *CreateCompanyUsecase) CreateCompany(req req.CreateCompanyRequest) error {
	// 이름 존재여부 확인
	userExists := cr.nameRepo.FindByName(req.Name)
	if !userExists {
		return fmt.Errorf("user not found")
	}

	// company 엔티티 생성
	time := time.Now()
	newCompany := entities.CompanyEntity{
		Id:          gocql.UUIDFromTime(time),
		Name:        req.Name,
		CompanyName: req.CompanyName,
		CreatedAt:   time,
	}
	return cr.companyRepo.CreateCompany(newCompany)
}
