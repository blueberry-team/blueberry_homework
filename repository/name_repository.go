package repository

type NameRepository struct {
	names []string
}

func NewNameRepository() *NameRepository { //얘는 함수 , func 함수명 (파라미터) 리턴타입{ }
	return &NameRepository{
		names: make([]string, 0),
	}
}

func (r *NameRepository) CreateName(name string) { //얘는 메서드 , func (리시버) 함수명 (파라미터) 리턴타입{ }. 리시버는 사실상 숨겨진 첫번째 파라미터!!!
	r.names = append(r.names, name)
}

func (r *NameRepository) GetName() []string {
	return r.names
}
