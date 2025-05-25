package user_usecase

import (
	"blueberry_homework/dto/request"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Login은 사용자 로그인을 처리합니다.
// false - 로그인 실패, true - 로그인 성공
func (u *UserUsecase) Login(req request.LoginRequest) (bool, error) {
	userExist, err := u.repo.FindByEmail(req.Email)
	if err != nil {
		return false, err
	}
	if !userExist {
		return false, errors.New("user not found")
	}

	hashedPassword, err := u.repo.GetHashedPassword(req.Email)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(req.Password))
	if err != nil {
		return false, errors.New("invalid password")
	}

	return true, nil
}
