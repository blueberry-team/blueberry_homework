package repository

// NameRepository는 이름을 관리하는 인터페이스입니다.
type NameRepository interface {
    // CreateName은 새로운 이름을 저장소에 추가합니다.
    CreateName(name string)
    // GetNames는 저장된 모든 이름을 반환합니다.
    GetNames() []string
}

// nameRepo는 NameRepository 인터페이스의 구현체입니다.
type nameRepo struct {
	// 슬라이스 선언
    names []string
}

// NewNameRepository는 새로운 NameRepository 인스턴스를 생성합니다.
// 초기화 함수 인 셈 => 생성자 함수
func NewNameRepository() NameRepository {
	// nameRepo 구조체의 포인터를 반환
    return &nameRepo{
        names: []string{}, // 슬라이스 초기화
    }
}

// CreateName은 새로운 이름을 저장소에 추가합니다.
func (r *nameRepo) CreateName(name string) {
    r.names = append(r.names, name)
}

// GetNames는 저장된 모든 이름을 반환합니다.
func (r *nameRepo) GetNames() []string {
    return r.names
}
