package repo_interface

import (
	"blueberry_homework/dto/response"
	"blueberry_homework/internal/domain/entities"

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

	// 해싱된 비밀번호 가져오는 함수
	GetHashedPassword(email string) (string, error)

	// 토큰 정보 가져오는 함수
	GetTokenInfo(email string) (string, string, error)

	// GetUser는 유저의 정보를 가져옵니다
	GetUser(id gocql.UUID) (response.UserResponse, error)

	// 사용자 정보 변경 함수
	ChangeUser(user entities.ChangeUserEntity) error
}
