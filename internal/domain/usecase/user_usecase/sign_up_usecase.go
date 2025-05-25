package user_usecase

import (
	"blueberry_homework/dto/request"
	"blueberry_homework/internal/domain/entities"
	"errors"
	"time"

	"github.com/gocql/gocql"
	"golang.org/x/crypto/bcrypt"
)

// SignUp은 사용자 생성을 처리합니다.
func (u *UserUsecase) SignUp(req request.SignUpRequest) error {
	// 이메일 중복 체크
	exists, err := u.repo.FindByEmail(req.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	time := time.Now()
	user := entities.UserEntity{
		Id:        gocql.UUIDFromTime(time),
		Email:     req.Email,
		Password:  string(hashedPassword),
		Name:      req.Name,
		Role:      req.Role,
		CreatedAt: time,
		UpdatedAt: time,
	}
	return u.repo.SignUp(user)
}
