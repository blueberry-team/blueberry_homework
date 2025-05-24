package usecase

import (
	"blueberry_homework/internal/domain/entities"
	"blueberry_homework/internal/domain/repo_interface"
	"blueberry_homework/dto/request"
	"blueberry_homework/dto/response"
	"errors"
	"time"

	"github.com/gocql/gocql"
)

type UserUsecase struct {
	repo repo_interface.UserRepository
}

func NewUserUsecase(r repo_interface.UserRepository) *UserUsecase {
	return &UserUsecase{
		repo: r,
	}
}

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
