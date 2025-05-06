package repository

import (
	"blueberry_homework/internal/domain/entities"
	"blueberry_homework/internal/domain/repo_interface"
	"blueberry_homework/internal/request"
	"sync"

	"fmt"
	"time"

	"github.com/gocql/gocql"
)

// nameRepo는 NameRepository 인터페이스의 구현체입니다.
type nameRepo struct {
	// 저장소
	session *gocql.Session
	// Mutex 추가
	mu sync.Mutex
}

// NewNameRepository는 새로운 NameRepository 인스턴스를 생성합니다.
// 초기화 함수 인 셈 => 생성자 함수
func NewNameRepository(session *gocql.Session) repo_interface.NameRepository {
	// nameRepo 구조체의 포인터를 반환
	return &nameRepo{
		session: session,
	}
}

// CreateName은 새로운 이름을 저장소에 추가합니다.
func (r *nameRepo) CreateName(entity entities.NameEntity) error {
	// Mutex로 락을 걸어 동시 접근을 방지합니다.
	r.mu.Lock()
	defer r.mu.Unlock()

	// 중복체크
	var existingName string
	err := r.session.Query(`
		SELECT name FROM names WHERE name = ? LIMIT 1
	`, entity.Name).Scan(&existingName)

	if err == nil {
		// 중복인 경우: 해당 name을 포함해 에러 메시지 리턴
		return fmt.Errorf("name already exists: %s", existingName)
	}
	if err != gocql.ErrNotFound {
		// 쿼리 오류
		return err
	}

	// INSERT
	return r.session.Query(`
		INSERT INTO names (id, name, created_at, updated_at)
		VALUES (?, ?, ?, ?)
	`, entity.Id, entity.Name, entity.CreatedAt, entity.UpdatedAt).Exec()
}

// GetNames는 저장된 모든 이름을 반환합니다.
func (r *nameRepo) GetNames() ([]entities.NameEntity, error) {
	// Iter() + Scan() 으로 여러 row
	iter := r.session.Query(`
		SELECT id, name, created_at, updated_at FROM names
	`).Iter()
	// 가져온 row 들을 저장
	var results []entities.NameEntity

	// iter.Scan()이 데이터를 복사할 변수들
	var id gocql.UUID
	var name string
	var createdAt, updatedAt time.Time

	for iter.Scan(&id, &name, &createdAt, &updatedAt) {
		results = append(results, entities.NameEntity{
			Id:        id,
			Name:      name,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}
	// iter가 끝났는데 에러가 존재한다면
	if err := iter.Close(); err != nil {
		return nil, fmt.Errorf("GetNames() query failed: %v", err)
	}

	return results, nil
}

// 이름으로 색인해서 삭제
func (r *nameRepo) DeleteByName(name string) error {
	// Mutex로 락을 걸어 동시 접근을 방지합니다.
	r.mu.Lock()
	defer r.mu.Unlock()

	// 1. 먼저 name으로 id 찾기
    var id gocql.UUID
    err := r.session.Query(`
        SELECT id FROM names WHERE name = ? LIMIT 1 ALLOW FILTERING
    `, name).Scan(&id)

    if err != nil {
        if err == gocql.ErrNotFound {
            return fmt.Errorf("name not found: %s", name)
        }
        return err
    }

    // 2. 찾은 id로 삭제
    return r.session.Query(`
        DELETE FROM names WHERE id = ?
    `, id).Exec()
}

// 여기서 time update 가 맞는 것인가 아닌것인가...
// 유저를 찾고나서 해야하니까 여기서 시간을 업데이트 하는 것이 옳은 것 같긴하다고 생각함
func (r *nameRepo) ChangeName(req request.ChangeNameRequest) error {
	// Mutex로 락을 걸어 동시 접근을 방지합니다.
	r.mu.Lock()
	defer r.mu.Unlock()

	// 이름 중복 체크 (선택적)
	var existingName string
	err := r.session.Query(`
		SELECT name FROM names WHERE name = ? LIMIT 1
	`, req.Name).Scan(&existingName)

	if err == nil {
		// 중복인 경우
		return fmt.Errorf("name already exists: %s", existingName)
	}
	if err != gocql.ErrNotFound {
		// 쿼리 오류
		return err
	}

	return r.session.Query(`
	UPDATE names SET name = ?, updated_at = ? WHERE id = ?
	`, req.Name, time.Now(), req.Id).Exec()
}

// 이름 존재여부 확인 함수
func (r *nameRepo) FindByName(name string) bool {
	var dummy string
	err := r.session.Query(`
		SELECT name FROM names WHERE name = ? LIMIT 1
	`, name).Scan(&dummy)

	if err == nil {
		return true // 존재함
	}
	if err == gocql.ErrNotFound {
		return false // 존재하지 않음
	}

	// 그 외의 쿼리 실패는 로그만 찍고 false 리턴
	fmt.Println("FindByName query error:", err)
	return false
}
