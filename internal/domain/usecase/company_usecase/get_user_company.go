package company_usecase

import (
	"blueberry_homework/dto/response"
	"errors"

	"github.com/gocql/gocql"
)

func (u *UserCompanyUsecase) GetUserCompany(userId string) (response.CompanyResponse, error) {
	parsedUserId, err := gocql.ParseUUID(userId)
	if err != nil {
		return response.CompanyResponse{}, err
	}

	userExist, err := u.userRepo.FindById(parsedUserId)
	if err != nil {
		return response.CompanyResponse{}, err
	}
	if !userExist {
		return response.CompanyResponse{}, errors.New("user not found")
	}

	companyExist, err := u.companyRepo.CheckCompanyWithUserId(parsedUserId)
	if err != nil {
		return response.CompanyResponse{}, err
	}
	if !companyExist {
		return response.CompanyResponse{}, errors.New("company not found")
	}

	return u.companyRepo.GetUserCompany(parsedUserId)
}
