package repo_interface

import (
	"blueberry_homework/dto/response"
	"blueberry_homework/internal/domain/entities"

	"github.com/gocql/gocql"
)

type CompanyRepository interface {
	CheckCompanyWithUserId(userId gocql.UUID) (bool, error)
	CreateCompany(company entities.CompanyEntity) error
	GetUserCompany(userId gocql.UUID) (response.CompanyResponse, error)
	ChangeCompany(changeCompany entities.CompanyEntity) error
	DeleteCompany(userId gocql.UUID) error
}
