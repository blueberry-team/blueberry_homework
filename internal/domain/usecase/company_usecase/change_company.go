package company_usecase

import (
	"blueberry_homework/dto/request"
	"blueberry_homework/internal/domain/entities"
	"errors"
	"time"

	"github.com/gocql/gocql"
)

func (u *CompanyUsecase) ChangeCompany(userId string, req request.ChangeCompanyRequest) error {
	parsedUserId, err := gocql.ParseUUID(userId)
	if err != nil {
		return err
	}
	companyExist, err := u.companyRepo.CheckCompanyWithUserId(parsedUserId)
	if err != nil {
		return err
	}
	if !companyExist {
		return errors.New("company not found")
	}

	now := time.Now()
	company := entities.ChangeCompanyEntity{
		UserId:         parsedUserId,
		CompanyName:    req.CompanyName,
		CompanyAddress: req.CompanyAddress,
		TotalStaff:     req.TotalStaff,
		UpdatedAt:      now,
	}
	return u.companyRepo.ChangeCompany(company)
}
