package entity

import "time"

// UserEntity 사용자 정보를 나타내는 구조체
type UserEntity struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}
