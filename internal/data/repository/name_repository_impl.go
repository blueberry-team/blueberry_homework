package repository

import (
	"blueberry_homework/internal/domain/entities"
	"blueberry_homework/internal/domain/repo_interface"
	"blueberry_homework/internal/dto"

	"fmt"
	"time"
)

// nameRepo는 NameRepository 인터페이스의 구현체입니다.
type nameRepo struct {
	// map 선언
	names []entities.NameEntity
}

// NewNameRepository는 새로운 NameRepository 인스턴스를 생성합니다.
// 초기화 함수 인 셈 => 생성자 함수
func NewNameRepository() repointerface.NameRepository {
	// nameRepo 구조체의 포인터를 반환
	return &nameRepo{
		names: []entities.NameEntity{},
	}
}

// CreateName은 새로운 이름을 저장소에 추가합니다.
func (r *nameRepo) CreateName(entity entities.NameEntity) error {
	for _, nameEntity := range r.names {
		if nameEntity.Name == entity.Name {
			return fmt.Errorf("name already exists")
		}
	}
	r.names = append(r.names, entity)
	return nil
}

// GetNames는 저장된 모든 이름을 반환합니다.
func (r *nameRepo) GetNames() []entities.NameEntity {
	return r.names
}

// DeleteName 은 인덱스에 해당하는 이름을 지우고 재정렬합니다.
func (r *nameRepo) DeleteByIndex(index int) {
	// 삭제 + 재정렬 (앞으로 당기기)
	// :index 는 인덱스 전까지, index: 는 인덱스에서부터 끝까지
	// The append built-in function appends elements to the end of a slice.
	// 즉 append(a, b) 라고 하면 a slice 뒤에 b slice 를 가져다 붙임
	// 그래서 한 개를 스킵할 수 있음
	r.names = append(r.names[:index], r.names[index+1:]...)
}

func (r *nameRepo) DeleteByName(name string) {
	// make([]T, 초기길이, 최대용량)
	// 새로운 슬라이스 선언이 아니라 make 를 쓰면 메모리 재할당없이 빠르게 추가가능
	filtered := make([]entities.NameEntity, 0, len(r.names))
	for _, item := range r.names {
		if item.Name != name {
			filtered = append(filtered, item)
		}
	}
	r.names = filtered
}

// 여기서 time update 가 맞는 것인가 아닌것인가...
// 유저를 찾고나서 해야하니까 여기서 시간을 업데이트 하는 것이 옳은 것 같긴하다고 생각함
func (r *nameRepo) ChangeName(req req.ChangeNameRequest) error {
	for i, name := range r.names {
		if name.Id == req.Id {
			r.names[i].Name = req.Name
			r.names[i].UpdatedAt = time.Now()
			return nil
		}
	}
	return fmt.Errorf("id not found")
}

// 이름 존재여부 확인 함수
func (r *nameRepo) FindByName(name string) (bool) {
	for _, entity := range r.names {
		if entity.Name == name {
			return true
		}
	}
	return false
}

// 

