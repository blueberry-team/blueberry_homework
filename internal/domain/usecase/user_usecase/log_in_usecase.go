package user_usecase

import (
	"blueberry_homework/dto/request"
	"blueberry_homework/dto/response"
	"blueberry_homework/utils/jwt"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Login은 사용자 로그인을 처리합니다.
// false - 로그인 실패, true - 로그인 성공
func (u *UserUsecase) Login(req request.LoginRequest) (response.TokenResponse, error) {
	userExist, err := u.repo.FindByEmail(req.Email)
	if err != nil {
		return response.TokenResponse{}, err
	}
	if !userExist {
		return response.TokenResponse{}, errors.New("user not found")
	}

	hashedPassword, err := u.repo.GetHashedPassword(req.Email)
	if err != nil {
		return response.TokenResponse{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	if err != nil {
		return response.TokenResponse{}, errors.New("invalid password")
	}

	userId, name, err := u.repo.GetTokenInfo(req.Email)
	if err != nil {
		return response.TokenResponse{}, err
	}

	token, err := jwt.GenerateToken(userId, req.Email, name)
	if err != nil {
		return response.TokenResponse{}, err
	}

	return response.TokenResponse{Token: token}, nil
}
