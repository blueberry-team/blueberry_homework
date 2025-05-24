package repository

import (
	"blueberry_homework_go_gin/entity"
	"time"
)

// NameRepository 이름을 저장하고 관리하는 저장소
type NameRepository struct {
	users []entity.UserEntity
}

// NewNameRepository 새로운 NameRepository 인스턴스를 생성
func NewNameRepository() *NameRepository {
	return &NameRepository{
		users: make([]entity.UserEntity, 0),
	}
}

// CreateName 새 사용자 이름을 추가
func (r *NameRepository) CreateName(user entity.UserEntity) {
	r.users = append(r.users, user)
}

// GetNames 모든 사용자 목록을 반환
func (r *NameRepository) GetNames() []entity.UserEntity {
	return r.users
}

// FindByName 이름으로 사용자를 찾음
func (r *NameRepository) FindByName(name string) *entity.UserEntity {
	for i, user := range r.users {
		if user.Name == name {
			return &r.users[i]
		}
	}
	return nil
}

// FindByID ID로 사용자를 찾음
func (r *NameRepository) FindByID(id string) *entity.UserEntity {
	for i, user := range r.users {
		if user.ID == id {
			return &r.users[i]
		}
	}
	return nil
}

// ChangeName 사용자 이름을 변경
func (r *NameRepository) ChangeName(id, newName string) bool {
	for i, user := range r.users {
		if user.ID == id {
			r.users[i].Name = newName
			r.users[i].UpdatedAt = time.Now()
			return true
		}
	}
	return false
}

// DeleteByIndex 특정 인덱스의 사용자를 삭제
func (r *NameRepository) DeleteByIndex(index int) bool {
	if index < 0 || index >= len(r.users) {
		return false
	}

	// 해당 인덱스 제거 (slice에서 해당 인덱스를 제외하고 다시 구성)
	r.users = append(r.users[:index], r.users[index+1:]...)
	return true
}

// DeleteByName 특정 이름을 가진 모든 사용자를 삭제
func (r *NameRepository) DeleteByName(name string) bool {
	originalLength := len(r.users)
	var filteredUsers []entity.UserEntity

	for _, user := range r.users {
		if user.Name != name {
			filteredUsers = append(filteredUsers, user)
		}
	}

	r.users = filteredUsers
	return len(r.users) < originalLength // 하나라도 삭제되었으면 true 반환
}
