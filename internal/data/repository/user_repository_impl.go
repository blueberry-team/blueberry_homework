package repository

import (
	"blueberry_homework/internal/domain/entities"
	"blueberry_homework/internal/domain/repo_interface"
	"blueberry_homework/dto/response"
	"sync"

	"fmt"

	"github.com/gocql/gocql"
)

// userRepo는 UserRepository 인터페이스의 구현체입니다.
type userRepo struct {
	// 저장소
	session *gocql.Session
	// Mutex 추가
	mu sync.Mutex
}

// NewUserRepository는 새로운 UserRepository 인스턴스를 생성합니다.
// 초기화 함수 인 셈 => 생성자 함수
func NewUserRepository(session *gocql.Session) repo_interface.UserRepository {
	// userRepo 구조체의 포인터를 반환
	return &userRepo{
		session: session,
	}
}

// 이메일로 사용자 찾기
func (r *userRepo) FindByEmail(email string) (bool, error) {
	var dummy gocql.UUID
	err := r.session.Query(`
		SELECT id FROM users WHERE email = ? LIMIT 1
	`, email).Scan(&dummy)

	if err != nil {
		if err == gocql.ErrNotFound {
			return false, nil // 사용자 없음 (에러가 아닌 정상 케이스로 처리)
		}
		return false, fmt.Errorf("FindByEmail query error: %v", err) // 그 외 쿼리 오류
	}
	return true, nil // 사용자 있음
}

// id로 사용자 찾기 함수
func (r *userRepo) FindById(id gocql.UUID) (bool, error) {
	var dummy gocql.UUID
	err := r.session.Query(`
		SELECT id FROM users WHERE id = ? LIMIT 1
	`, id).Scan(&dummy)

	if err != nil {
		if err == gocql.ErrNotFound {
			return false, nil // 사용자 없음
		}
		return false, fmt.Errorf("FindById query error: %v", err) // 그 외 쿼리 오류
	}
	return true, nil // 사용자 있음
}

// 유저의 역할을 확인하는 함수
func (r *userRepo) CheckRole(id gocql.UUID) (string, error) {
	var role string
	err := r.session.Query(`
		SELECT role FROM users WHERE id = ? LIMIT 1
	`, id).Scan(&role)

	if err != nil {
		if err == gocql.ErrNotFound {
			return "", fmt.Errorf("user not found with id: %s", id.String())
		}
		return "", fmt.Errorf("failed to check role: %v", err)
	}
	return role, nil
}

// SignUp은 새로운 사용자를 저장소에 추가합니다.
func (r *userRepo) SignUp(entity entities.UserEntity) error {
	// Mutex로 락을 걸어 동시 접근을 방지합니다.
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.session.Query(`
		INSERT INTO users (id, email, password, name, role, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, entity.Id, entity.Email, entity.Password, entity.Name, entity.Role, entity.CreatedAt, entity.UpdatedAt).Exec()
}

// 로그인 함수 -> GetHashedPassword 함수로 변경
// 로그인이라는 기능에 있어서 repo 가 해야할 일은 기존의 hashed password 를 가져오는 것뿐이기 때문에
// 따라서 로그인 함수가 아닌 GetHashedPassword 함수가 됩니다.
func (r *userRepo) GetHashedPassword(email string) (string, error) {
	var hashed_password string
	err := r.session.Query(`
		SELECT password FROM users WHERE email = ? LIMIT 1
	`, email).Scan(&hashed_password)

	if err != nil {
		if err == gocql.ErrNotFound {
			return "", nil // 사용자 없음
		}
		return "", fmt.Errorf("login query error: %v", err)
	}

	return hashed_password, nil
}

// GetUser는 유저의 정보를 가져옵니다
func (r *userRepo) GetUser(id gocql.UUID) (response.UserResponse, error) {
	var user entities.UserEntity
	err := r.session.Query(`
		SELECT email, name, role, created_at, updated_at FROM users WHERE id = ? LIMIT 1
	`, id).Scan(&user.Email, &user.Name, &user.Role, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == gocql.ErrNotFound {
			return response.UserResponse{}, fmt.Errorf("user not found with id: %s", id.String())
		}
		return response.UserResponse{}, fmt.Errorf("failed to get user: %v", err)
	}

	return response.UserResponse{
		Email:     user.Email,
		Name:      user.Name,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// 사용자 정보 변경 함수
func (r *userRepo) ChangeUser(user entities.UserEntity) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.session.Query(`
		UPDATE users
		SET name = ?, role = ?, updated_at = ?
		WHERE id = ?
	`, user.Name, user.Role, user.UpdatedAt, user.Id).Exec()
}
