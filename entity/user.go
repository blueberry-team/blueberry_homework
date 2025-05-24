
package entity

import (
	"time"

	"github.com/google/uuid"
)

// UserEntity 사용자 정보를 나타내는 구조체
type UserEntity struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// NewUserEntity 새로운 UserEntity 생성
func NewUserEntity(name string) UserEntity {
	now := time.Now()
	return UserEntity{
		ID:        uuid.New().String(),
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
