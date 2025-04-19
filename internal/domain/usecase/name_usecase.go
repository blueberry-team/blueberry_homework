package usecase

import (
	"blueberry_homework/internal/domain/entities"
	"blueberry_homework/internal/domain/repo_interface"
	"time"
)

type NameUsecase struct {
	repo repointerface.NameRepository
}

func NewNameUsecase(r repointerface.NameRepository) *NameUsecase {
	return &NameUsecase{
		repo: r,
	}
}

func (u *NameUsecase) CreateName(name string) error {
	entity := entities.NameEntity{
		Name:      name,
		CreatedAt: time.Now(),
	}
	u.repo.CreateName(entity)
	return nil
}

// GetNames는 저장된 모든 이름 리스트를 반환합니다.
func (u *NameUsecase) GetNames() []entities.NameEntity {
	return u.repo.GetNames()
}

// DeleteByIndex는 인덱스로 이름을 삭제합니다.
func (u *NameUsecase) DeleteByIndex(index int) {
	u.repo.DeleteByIndex(index)
}

// DeleteByName은 이름이 일치하는 모든 항목을 삭제합니다.
func (u *NameUsecase) DeleteByName(name string) {
	u.repo.DeleteByName(name)
}
