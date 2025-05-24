package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"blueberry_homework_go_gin/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client   *mongo.Client
	Database *mongo.Database
)

// InitMongoDB MongoDB 연결을 초기화
func InitMongoDB(cfg *config.Config) error {
	// MongoDB 클라이언트 옵션 설정
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)

	// 연결 타임아웃 설정
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// MongoDB에 연결
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("MongoDB 연결 실패: %v", err)
	}

	// 연결 테스트
	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("MongoDB 핑 실패: %v", err)
	}

	Client = client
	Database = client.Database(cfg.MongoDatabase)

	log.Printf("✅ MongoDB 연결 성공: %s/%s", cfg.MongoURI, cfg.MongoDatabase)

	// 컬렉션 초기화
	if err := initCollections(); err != nil {
		return fmt.Errorf("컬렉션 초기화 실패: %v", err)
	}

	return nil
}

// initCollections 필요한 컬렉션들을 초기화
func initCollections() error {
	ctx := context.Background()

	// users 컬렉션 생성 (기존 사용자 - 호환성 유지)
	err := Database.CreateCollection(ctx, "users")
	if err != nil {
		log.Printf("users 컬렉션: %v", err)
	}

	// auth_users 컬렉션 생성 (인증 사용자)
	err = Database.CreateCollection(ctx, "auth_users")
	if err != nil {
		log.Printf("auth_users 컬렉션: %v", err)
	}

	// companies 컬렉션 생성
	err = Database.CreateCollection(ctx, "companies")
	if err != nil {
		log.Printf("companies 컬렉션: %v", err)
	}

	log.Println("✅ 컬렉션 초기화 완료")
	return nil
}

// CloseMongoDB MongoDB 연결을 종료
func CloseMongoDB() error {
	if Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := Client.Disconnect(ctx); err != nil {
			return fmt.Errorf("MongoDB 연결 종료 실패: %v", err)
		}
		log.Println("✅ MongoDB 연결 종료")
	}
	return nil
}

// GetCollection 지정된 컬렉션을 반환
func GetCollection(name string) *mongo.Collection {
	return Database.Collection(name)
}
