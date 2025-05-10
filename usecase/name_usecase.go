package usecase

import (
	"blueberry_homework_go_gin/entity"
	"blueberry_homework_go_gin/repository"
	"time"
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
func (u *NameUseCase) CreateName(name string) {
	user := entity.UserEntity{
		Name:      name,
		CreatedAt: time.Now(),
	}
	u.repo.CreateName(user)
}

// GetNames 모든 이름 목록을 조회
func (u *NameUseCase) GetNames() []entity.UserEntity {
	return u.repo.GetNames()
}

// DeleteByIndex 특정 인덱스의 이름을 삭제
func (u *NameUseCase) DeleteByIndex(index int) bool {
	return u.repo.DeleteByIndex(index)
}

// DeleteByName 특정 이름을 가진 모든 항목을 삭제
func (u *NameUseCase) DeleteByName(name string) bool {
	return u.repo.DeleteByName(name)
}
