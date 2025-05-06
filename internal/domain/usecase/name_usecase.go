package usecase

import (
	"blueberry_homework/internal/domain/entities"
	"blueberry_homework/internal/domain/repo_interface"
	"blueberry_homework/internal/request"
	"time"

	"github.com/gocql/gocql"
)

type NameUsecase struct {
	repo repo_interface.NameRepository
}

func NewNameUsecase(r repo_interface.NameRepository) *NameUsecase {
	return &NameUsecase{
		repo: r,
	}
}

// 이름 생성함수
// createdAt, updatedAt 초기화도 여기서 이뤄집니다.
func (u *NameUsecase) CreateName(name string) error {
	time := time.Now()
	entity := entities.NameEntity{
		Id:        gocql.UUIDFromTime(time),
		Name:      name,
		CreatedAt: time,
		UpdatedAt: time,
	}
	return u.repo.CreateName(entity)
}

// GetNames는 저장된 모든 이름 리스트를 반환합니다.
func (u *NameUsecase) GetNames() ([]entities.NameEntity, error) {
	return u.repo.GetNames()
}

// DeleteByName은 이름이 일치하는 모든 항목을 삭제합니다.
func (u *NameUsecase) DeleteByName(name string) error {
	return u.repo.DeleteByName(name)
}

// changeName은 id 로 유저를 검색해서 이름을 변경합니다.
func (u *NameUsecase) ChangeName(req request.ChangeNameRequest) error {
	return u.repo.ChangeName(req)
}
