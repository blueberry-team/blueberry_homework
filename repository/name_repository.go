package repository

type NameRepository struct {
	names []map[string]string
}

func NewNameRepository() *NameRepository { //얘는 함수 , func 함수명 (파라미터) 리턴타입{ }
	return &NameRepository{
		names: make([]map[string]string, 0),
	}
}

func (r *NameRepository) CreateName(name string) { //얘는 메서드 , func (리시버) 함수명 (파라미터) 리턴타입{ }. 리시버는 사실상 숨겨진 첫번째 파라미터!!!
	r.names = append(r.names, map[string]string{"name": name})
}

func (r *NameRepository) GetName() []map[string]string {
	return r.names
}

func (r *NameRepository) DeleteName(index int) bool {
	if index < 0 || index >= len(r.names) {
		return false
	}
	r.names = append(r.names[:index], r.names[index+1:]...)
	return true
}
