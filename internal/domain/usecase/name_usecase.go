package usecase

import (
	"blueberry_homework/internal/domain/entities"
	"blueberry_homework/internal/domain/repo_interface"
	req "blueberry_homework/internal/dto"
	"time"

	"github.com/google/uuid"
)

type NameUsecase struct {
	repo repointerface.NameRepository
}

func NewNameUsecase(r repointerface.NameRepository) *NameUsecase {
	return &NameUsecase{
		repo: r,
	}
}

// 이름 생성함수
// createdAt, updatedAt 초기화도 여기서 이뤄집니다.
func (u *NameUsecase) CreateName(name string) error {
	time := time.Now()
	entity := entities.NameEntity{
		Id: uuid.New().String(),
		Name:      name,
		CreatedAt: time,
		UpdatedAt: time,
	}
	return u.repo.CreateName(entity)
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

// changeName은 id 로 유저를 검색해서 이름을 변경합니다.
func (u *NameUsecase) ChangeName(req req.ChangeNameRequest) error {
	return u.repo.ChangeName(req)
}
