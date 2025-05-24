package repository

import (
	"context"
	"time"

	"blueberry_homework_go_gin/db"
	"blueberry_homework_go_gin/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const usersCollection = "users"

// NameRepository 이름을 저장하고 관리하는 저장소
type NameRepository struct {
	collection *mongo.Collection
}

// NewNameRepository 새로운 NameRepository 인스턴스를 생성
func NewNameRepository() *NameRepository {
	return &NameRepository{
		collection: db.GetCollection(usersCollection),
	}
}

// CreateName 새 사용자 이름을 추가
func (r *NameRepository) CreateName(user entity.UserEntity) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, user)
	return err
}

// GetNames 모든 사용자 목록을 반환
func (r *NameRepository) GetNames() ([]entity.UserEntity, error) {
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

// FindByName 이름으로 사용자를 찾음
func (r *NameRepository) FindByName(name string) (*entity.UserEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user entity.UserEntity
	filter := bson.D{{Key: "name", Value: name}}

	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // 문서가 없으면 nil 반환
		}
		return nil, err
	}

	return &user, nil
}

// FindByID ID로 사용자를 찾음
func (r *NameRepository) FindByID(id string) (*entity.UserEntity, error) {
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

// ChangeName 사용자 이름을 변경
func (r *NameRepository) ChangeName(id, newName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{Key: "id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: newName},
			{Key: "updatedAt", Value: time.Now()},
		}},
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

// DeleteByIndex 특정 인덱스의 사용자를 삭제
func (r *NameRepository) DeleteByIndex(index int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 모든 사용자를 가져와서 인덱스에 해당하는 사용자를 찾음
	users, err := r.GetNames()
	if err != nil {
		return err
	}

	if index < 0 || index >= len(users) {
		return mongo.ErrNoDocuments
	}

	// 해당 인덱스의 사용자 ID로 삭제
	userToDelete := users[index]
	filter := bson.D{{Key: "id", Value: userToDelete.ID}}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

// DeleteByName 특정 이름을 가진 모든 사용자를 삭제
func (r *NameRepository) DeleteByName(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{Key: "name", Value: name}}
	result, err := r.collection.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
