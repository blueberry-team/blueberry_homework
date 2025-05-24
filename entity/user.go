package entity

import (
	"time"

	"github.com/google/uuid"
)

// UserRole 사용자 역할 타입
type UserRole string

const (
	RoleBoss   UserRole = "boss"
	RoleWorker UserRole = "worker"
)

// UserEntity 사용자 정보를 나타내는 구조체 (인증 시스템 포함)
type UserEntity struct {
	ID        string    `json:"id" bson:"id"`
	Email     string    `json:"email" bson:"email"`
	Password  string    `json:"-" bson:"password"` // JSON 응답에서 제외
	Name      string    `json:"name" bson:"name"`
	Role      UserRole  `json:"role" bson:"role"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

// SimpleUserEntity 기존 시스템용 단순 사용자 구조체 (호환성 유지)
type SimpleUserEntity struct {
	ID        string    `json:"id" bson:"id"`
	Name      string    `json:"name" bson:"name"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

// NewUserEntity 새로운 UserEntity 생성 (인증 시스템용)
func NewUserEntity(email, hashedPassword, name string, role UserRole) UserEntity {
	now := time.Now()
	return UserEntity{
		ID:        uuid.New().String(),
		Email:     email,
		Password:  hashedPassword,
		Name:      name,
		Role:      role,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// NewSimpleUserEntity 새로운 SimpleUserEntity 생성 (기존 시스템용)
func NewSimpleUserEntity(name string) SimpleUserEntity {
	now := time.Now()
	return SimpleUserEntity{
		ID:        uuid.New().String(),
		Name:      name,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// IsValidRole 유효한 역할인지 확인
func IsValidRole(role string) bool {
	return role == string(RoleBoss) || role == string(RoleWorker)
}

// IsBoss 보스 역할인지 확인
func (u *UserEntity) IsBoss() bool {
	return u.Role == RoleBoss
}

// UserPublicInfo 공개 사용자 정보 (비밀번호 제외)
type UserPublicInfo struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Role      UserRole  `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ToPublicInfo 공개 정보로 변환
func (u *UserEntity) ToPublicInfo() UserPublicInfo {
	return UserPublicInfo{
		ID:        u.ID,
		Email:     u.Email,
		Name:      u.Name,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
