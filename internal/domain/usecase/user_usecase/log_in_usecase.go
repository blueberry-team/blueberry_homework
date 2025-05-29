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
	userId, err := u.repo.FindByEmail(req.Email)
	if err != nil {
		return response.TokenResponse{}, err
	}

	hashedPassword, err := u.repo.GetHashedPassword(req.Email)
	if err != nil {
		return response.TokenResponse{}, err
	}

	// 비번 검증
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	if err != nil {
		return response.TokenResponse{}, errors.New("invalid password")
	}

	email, name, err := u.repo.GetTokenInfo(userId)
	if err != nil {
		return response.TokenResponse{}, err
	}

	token, err := jwt.GenerateToken(userId.String(), email, name)
	if err != nil {
		return response.TokenResponse{}, err
	}

	return response.TokenResponse{Token: token}, nil
}
