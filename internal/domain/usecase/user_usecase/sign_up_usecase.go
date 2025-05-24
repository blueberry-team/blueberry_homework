package user_usecase

import (
	"blueberry_homework/internal/domain/entities"
	"errors"
	"time"

	"github.com/gocql/gocql"
)

// SignUp은 사용자 생성을 처리합니다.
// createdAt, updatedAt 초기화도 여기서 이뤄집니다.
func (u *UserUsecase) SignUp(email string, password string, name string, role string) error {
	// 이메일 중복 체크
	exists, err := u.repo.FindByEmail(email)
	if err != nil {
		return err // FindByEmail에서 오류 발생 시 반환
	}
	if exists {
		return errors.New("email already exists") // 사용자 정의 오류 또는 특정 오류 타입 사용 가능
	}

	// TODO: 비밀번호 해싱 로직 필요

	time := time.Now()
	user := entities.UserEntity{
		Id:        gocql.UUIDFromTime(time),
		Email:     email,
		Password:  password, // 실제 프로덕션에서는 해싱 처리가 필요합니다.
		Name:      name,
		Role:      role,
		CreatedAt: time,
		UpdatedAt: time,
	}
	return u.repo.SignUp(user)
}
