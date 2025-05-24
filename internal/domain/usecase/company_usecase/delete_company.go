package company_usecase

import (
	"errors"

	"github.com/gocql/gocql"
)

func (u *UserCompanyUsecase) DeleteCompany(userId string) error {
	parsedUserId, err := gocql.ParseUUID(userId)
	if err != nil {
		return err
	}

	userExist, err := u.userRepo.FindById(parsedUserId)
	if err != nil {
		return err
	}
	if !userExist {
		return errors.New("user not found")
	}

	companyExist, err := u.companyRepo.CheckCompanyWithUserId(parsedUserId)
	if err != nil {
		return err
	}
	if !companyExist {
		return errors.New("company not found")
	}
	
	return u.companyRepo.DeleteCompany(parsedUserId)
}
