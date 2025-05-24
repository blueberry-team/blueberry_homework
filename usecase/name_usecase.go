package usecase

import (
	"errors"

	"blueberry_homework_go_gin/entity"
	"blueberry_homework_go_gin/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

// NameUseCase 이름 관련 비즈니스 로직을 담당하는 유스케이스
type NameUseCase struct {
	nameRepo *repository.NameRepository
}

// NewNameUseCase 새로운 NameUseCase 인스턴스를 생성
func NewNameUseCase(nameRepo *repository.NameRepository) *NameUseCase {
	return &NameUseCase{
		nameRepo: nameRepo,
	}
}

// CreateName 새로운 이름을 생성하고 저장
func (u *NameUseCase) CreateName(name string) error {
	// 중복 이름 체크
	existingUser, err := u.nameRepo.FindByName(name)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("A name with the same value already exists")
	}

	// 새 사용자 엔티티 생성
	user := entity.NewUserEntity(name)
	return u.nameRepo.CreateName(user)
}

// GetNames 모든 이름 목록을 조회
func (u *NameUseCase) GetNames() ([]entity.UserEntity, error) {
	return u.nameRepo.GetNames()
}

// ChangeName 사용자 이름을 변경
func (u *NameUseCase) ChangeName(id, newName string) error {
	// 사용자 존재 여부 확인
	user, err := u.nameRepo.FindByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("User not found")
	}

	// 이름이 동일한 경우 (변경 없음)
	if user.Name == newName {
		return errors.New("A name with the same value already exists.")
	}

	// 다른 사용자와 이름 중복 체크
	existingUser, err := u.nameRepo.FindByName(newName)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("A name with the same value already exists.")
	}

	// 이름 변경
	return u.nameRepo.ChangeName(id, newName)
}

// FindByName 이름으로 사용자 찾기 (내부 로직용)
func (u *NameUseCase) FindByName(name string) (*entity.UserEntity, error) {
	return u.nameRepo.FindByName(name)
}

// DeleteByIndex 특정 인덱스의 이름을 삭제
func (u *NameUseCase) DeleteByIndex(index int) error {
	err := u.nameRepo.DeleteByIndex(index)
	if err == mongo.ErrNoDocuments {
		return errors.New("index out of range")
	}
	return err
}

// DeleteByName 특정 이름을 가진 모든 항목을 삭제
func (u *NameUseCase) DeleteByName(name string) error {
	err := u.nameRepo.DeleteByName(name)
	if err == mongo.ErrNoDocuments {
		return errors.New("name not found")
	}
	return err
}
