package company_usecase

import (
	"blueberry_homework/dto/request"
	"blueberry_homework/internal/domain/entities"
	"errors"
	"time"

	"github.com/gocql/gocql"
)

func (u *UserCompanyUsecase) CreateCompany(createCompanyRequest request.CreateCompanyRequest) error {
	userId, err := gocql.ParseUUID(createCompanyRequest.UserID)
	if err != nil {
		return err
	}

	// 유저 존재 확인
    userExist, err := u.userRepo.FindById(userId)
    if err != nil {
        return err
    }
    if !userExist {
        return errors.New("user not found")
    }

	time := time.Now()
	entity := entities.CompanyEntity{
		Id:             gocql.UUIDFromTime(time),
		UserId:         userId,
		CompanyName:    createCompanyRequest.CompanyName,
		CompanyAddress: createCompanyRequest.CompanyAddress,
		TotalStaff:     createCompanyRequest.TotalStaff,
		CreatedAt:      time,
		UpdatedAt:      time,
	}
	return u.companyRepo.CreateCompany(entity)
}
