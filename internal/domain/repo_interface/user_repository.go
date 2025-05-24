package repo_interface

import (
	"blueberry_homework/internal/domain/entities"
	"blueberry_homework/dto/response"

	"github.com/gocql/gocql"
)

// UserRepository는 사용자를 관리하는 인터페이스입니다.
type UserRepository interface {
	// 이메일로 사용자 찾기 함수
	FindByEmail(email string) (bool, error)

	// id로 사용자 찾기 함수
	FindById(id gocql.UUID) (bool, error)

	// 유저의 역할을 확인하는 함수
	CheckRole(id gocql.UUID) (string, error)

	// 회원가입 함수
	SignUp(entity entities.UserEntity) error

	// 로그인 처리 함수
	Login(email string, password string) (bool, error)

	// GetUser는 유저의 정보를 가져옵니다
	GetUser(id gocql.UUID) (response.UserResponse, error)

	// 사용자 정보 변경 함수
	ChangeUser(user entities.UserEntity) error
}
