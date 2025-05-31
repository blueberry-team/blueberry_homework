package auth_usecase

import (
	"blueberry_homework/dto/response"
	"blueberry_homework/internal/domain/repo_interface"
	"blueberry_homework/utils/jwt"

	"github.com/gocql/gocql"
)

type AuthUsecase struct {
	userRepo repo_interface.UserRepository
}

func NewAuthUsecase(userRepo repo_interface.UserRepository) *AuthUsecase {
	return &AuthUsecase{userRepo: userRepo}
}

func (u *AuthUsecase) RefreshToken(userId string) (response.TokenResponse, error) {
	parsedId, err := gocql.ParseUUID(userId)
	if err != nil {
		return response.TokenResponse{}, err
	}

	email, name, err := u.userRepo.GetTokenInfo(parsedId)
	if err != nil {
		return response.TokenResponse{}, err
	}

	token, err := jwt.GenerateToken(userId, email, name)
	if err != nil {
		return response.TokenResponse{}, err
	}

	return response.TokenResponse{Token: token}, nil
}
