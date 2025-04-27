package usecase

import (
	"blueberry_homework/internal/domain/entities"
	"blueberry_homework/internal/domain/repo_interface"
	"blueberry_homework/internal/dto"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type CreateCompanyUsecase struct {
	nameRepo repointerface.NameRepository
	companyRepo repointerface.CompanyRepository
}

func NewCreateCompanyUsecase (n repointerface.NameRepository, c repointerface.CompanyRepository) *CreateCompanyUsecase {
	return &CreateCompanyUsecase{
		nameRepo: n,
		companyRepo: c,
	}
}

func (cr *CreateCompanyUsecase) CreateCompany (req req.CreateCompanyRequest) error {
	// 이름 존재여부 확인
	userExists := cr.nameRepo.FindByName(req.Name)
	if !userExists {
		return fmt.Errorf("user not found")
	}

	// company 엔티티 생성
	now := time.Now()
	newCompany := entities.CompanyEntity{
		Id:          uuid.New().String(),
		Name:        req.Name,
		CompanyName: req.CompanyName,
		CreatedAt:   now,
	}
	return cr.companyRepo.CreateCompany(newCompany)
}
