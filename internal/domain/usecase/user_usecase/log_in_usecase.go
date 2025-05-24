package user_usecase

import (
	"errors"
)

// Login은 사용자 로그인을 처리합니다.
// false - 로그인 실패, true - 로그인 성공
func (u *UserUsecase) Login(email string, password string) (bool, error) {
	userExist, err := u.repo.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if !userExist {
		return false, errors.New("user not found")
	}

	// TODO: 비밀번호 해싱 및 확인 로직 필요

	return u.repo.Login(email, password)
}
