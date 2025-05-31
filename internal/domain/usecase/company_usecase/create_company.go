package company_usecase

import (
	"blueberry_homework/dto/request"
	"blueberry_homework/internal/domain/entities"
	"errors"
	"time"

	"github.com/gocql/gocql"
)

func (u *CompanyUsecase) CreateCompany(userId string, createCompanyRequest request.CreateCompanyRequest) error {
	parsedId, err := gocql.ParseUUID(userId)
	if err != nil {
		return err
	}

	// 유저 존재 확인
    userExist, err := u.userRepo.FindById(parsedId)
    if err != nil {
        return err
    }
    if !userExist {
        return errors.New("user not found")
    }

	// 회사 존재 확인
	companyExist, err := u.companyRepo.CheckCompanyWithUserId(parsedId)
	if err != nil {
		return err
	}
	if companyExist {
		return errors.New("company already exists")
	}

	now := time.Now()
	entity := entities.CompanyEntity{
		Id:             gocql.UUIDFromTime(now),
		UserId:         parsedId,
		CompanyName:    createCompanyRequest.CompanyName,
		CompanyAddress: createCompanyRequest.CompanyAddress,
		TotalStaff:     createCompanyRequest.TotalStaff,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	return u.companyRepo.CreateCompany(entity)
}
