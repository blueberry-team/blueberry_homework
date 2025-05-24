package user_usecase

import (
	"blueberry_homework/dto/response"
	"errors"

	"github.com/gocql/gocql"
)

// GetUser는 ID로 특정 사용자 정보를 가져옵니다.
func (u *UserUsecase) GetUser(id string) (response.UserResponse, error) {
	// 유저 ID UUID 변환
	parsedId, err := gocql.ParseUUID(id)
	if err != nil {
		return response.UserResponse{}, err
	}

	userExist, err := u.repo.FindById(parsedId)
	if err != nil {
		return response.UserResponse{}, err
	}
	if !userExist {
		return response.UserResponse{}, errors.New("user not found")
	}

	return u.repo.GetUser(parsedId)
}
