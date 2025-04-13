package repository

import (
	"blueberry_homework/internal/models"
)

// NameRepository는 이름을 관리하는 인터페이스입니다.
type NameRepository interface {
    // CreateName은 새로운 이름을 저장소에 추가합니다.
    CreateName(name string)
    // GetNames는 저장된 모든 이름을 반환합니다.
    GetNames() []models.NameModel
    // DeleteName 은 index 를 받아 해당하는 이름을 삭제하고 재정렬합니다
    DeleteName(index int)
}

// nameRepo는 NameRepository 인터페이스의 구현체입니다.
type nameRepo struct {
	// map 선언
    names map[string][]models.NameModel
}

// NewNameRepository는 새로운 NameRepository 인스턴스를 생성합니다.
// 초기화 함수 인 셈 => 생성자 함수
func NewNameRepository() NameRepository {
	// nameRepo 구조체의 포인터를 반환
    return &nameRepo{
        names: map[string][]models.NameModel{
            "names": {}, // 슬라이스 초기화
        },
    }
}

// CreateName은 새로운 이름을 저장소에 추가합니다.
func (r *nameRepo) CreateName(name string) {
    r.names["names"] = append(r.names["names"], models.NameModel{Name:name})
}

// GetNames는 저장된 모든 이름을 반환합니다.
func (r *nameRepo) GetNames() []models.NameModel {
    return r.names["names"]
}

// DeleteName 은 인덱스에 해당하는 이름을 지우고 재정렬합니다.
func (r *nameRepo) DeleteName(index int) () {
	// 삭제 + 재정렬 (앞으로 당기기)
    list := r.names["names"]
    // :index 는 인덱스 전까지, index: 는 인덱스에서부터 끝까지
    // The append built-in function appends elements to the end of a slice.
    // 즉 append(a, b) 라고 하면 a slice 뒤에 b slice 를 가져다 붙임
    // 그래서 한 개를 스킵할 수 있음
	list = append(list[:index], list[index+1:]...)
	r.names["names"] = list
}

