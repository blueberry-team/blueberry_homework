package user_usecase

import (
	"blueberry_homework/dto/request"
	"blueberry_homework/internal/domain/entities"
	"errors"
	"time"

	"github.com/gocql/gocql"
)

// ChangeUser는 사용자 정보를 변경합니다.
func (u *UserUsecase) ChangeUser(req request.ChangeUserRequest) error {
	parsedId, err := gocql.ParseUUID(req.Id)
	if err != nil {
		return err
	}
	userExist, err := u.repo.FindById(parsedId)
	if err != nil {
		return err
	}
	if !userExist {
		return errors.New("user not found")
	}
	entity := entities.UserEntity{
		Id:        parsedId,
		Name:      req.Name,
		Role:      req.Role,
		UpdatedAt: time.Now(),
	}
	return u.repo.ChangeUser(entity)
}
