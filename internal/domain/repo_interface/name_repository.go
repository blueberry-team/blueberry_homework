package repointerface

import "blueberry_homework/internal/domain/entities"

// NameRepository는 이름을 관리하는 인터페이스입니다.
type NameRepository interface {
	// CreateName은 새로운 이름을 저장소에 추가합니다.
	CreateName(entity entities.NameEntity)

	// GetNames는 저장된 모든 이름을 반환합니다.
	GetNames() []entities.NameEntity

	// DeleteByIndex 은 index 를 받아 해당하는 이름을 삭제하고 재정렬합니다
	DeleteByIndex(index int)

	// DeleteByName 은 이름을 받아 해당하는 이름을 삭제하고 재정렬합니다
	DeleteByName(name string)
}
