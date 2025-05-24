package repository

import (
	"context"
	"time"

	"blueberry_homework_go_gin/db"
	"blueberry_homework_go_gin/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const authUsersCollection = "auth_users"

// AuthRepository 인증 관련 사용자를 저장하고 관리하는 저장소
type AuthRepository struct {
	collection *mongo.Collection
}

// NewAuthRepository 새로운 AuthRepository 인스턴스를 생성
func NewAuthRepository() *AuthRepository {
	return &AuthRepository{
		collection: db.GetCollection(authUsersCollection),
	}
}

// CreateUser 새 사용자를 추가 (회원가입)
func (r *AuthRepository) CreateUser(user entity.UserEntity) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, user)
	return err
}

// FindUserByEmail 이메일로 사용자를 찾음
func (r *AuthRepository) FindUserByEmail(email string) (*entity.UserEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user entity.UserEntity
	filter := bson.D{{Key: "email", Value: email}}

	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // 문서가 없으면 nil 반환
		}
		return nil, err
	}

	return &user, nil
}

// FindUserByID ID로 사용자를 찾음
func (r *AuthRepository) FindUserByID(id string) (*entity.UserEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user entity.UserEntity
	filter := bson.D{{Key: "id", Value: id}}

	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // 문서가 없으면 nil 반환
		}
		return nil, err
	}

	return &user, nil
}

// UpdateUser 사용자 정보를 업데이트
func (r *AuthRepository) UpdateUser(id string, updates bson.D) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{Key: "id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: append(updates, bson.E{Key: "updatedAt", Value: time.Now()})},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

// GetAllUsers 모든 사용자 목록을 반환 (관리자용)
func (r *AuthRepository) GetAllUsers() ([]entity.UserEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []entity.UserEntity
	if err = cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	// nil 대신 빈 슬라이스 반환
	if users == nil {
		users = []entity.UserEntity{}
	}

	return users, nil
}
