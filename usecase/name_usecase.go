package usecase

import (
	"errors"
	"blueberry_homework_go_gin/entity"
	"blueberry_homework_go_gin/repository"
)

// NameUseCase 이름 관련 비즈니스 로직을 담당하는 유스케이스
type NameUseCase struct {
	repo *repository.NameRepository
}

// NewNameUseCase 새로운 NameUseCase 인스턴스를 생성
func NewNameUseCase(repo *repository.NameRepository) *NameUseCase {
	return &NameUseCase{
		repo: repo,
	}
}

// CreateName 새로운 이름을 생성하고 저장
func (u *NameUseCase) CreateName(name string) error {
	// 중복 이름 체크
	if u.repo.FindByName(name) != nil {
		return errors.New("A name with the same value already exists")
	}

	user := entity.NewUserEntity(name)
	u.repo.CreateName(user)
	return nil
}

// ChangeName 사용자 이름을 변경
func (u *NameUseCase) ChangeName(id, newName string) error {
	// 사용자 찾기
	user := u.repo.FindByID(id)
	if user == nil {
		return errors.New("User not found")
	}

	// 이름이 동일한 경우 (변경 없음)
	if user.Name == newName {
		return errors.New("A name with the same value already exists.")
	}

	// 다른 사용자와 이름 중복 체크
	existingUser := u.repo.FindByName(newName)
	if existingUser != nil {
		return errors.New("A name with the same value already exists.")
	}

	// 이름 변경
	u.repo.ChangeName(id, newName)
	return nil
}

// GetNames 모든 이름 목록을 조회
func (u *NameUseCase) GetNames() []entity.UserEntity {
	return u.repo.GetNames()
}

// FindByName 이름으로 사용자 찾기
func (u *NameUseCase) FindByName(name string) *entity.UserEntity {
	return u.repo.FindByName(name)
}

// DeleteByIndex 특정 인덱스의 이름을 삭제
func (u *NameUseCase) DeleteByIndex(index int) bool {
	return u.repo.DeleteByIndex(index)
}

// DeleteByName 특정 이름을 가진 모든 항목을 삭제
func (u *NameUseCase) DeleteByName(name string) bool {
	return u.repo.DeleteByName(name)
}
