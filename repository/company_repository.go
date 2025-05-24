package repository

import (
	"context"
	"time"

	"blueberry_homework_go_gin/db"
	"blueberry_homework_go_gin/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const companiesCollection = "companies"

// CompanyRepository 회사 정보를 저장하고 관리하는 저장소
type CompanyRepository struct {
	collection *mongo.Collection
}

// NewCompanyRepository 새로운 CompanyRepository 인스턴스를 생성
func NewCompanyRepository() *CompanyRepository {
	return &CompanyRepository{
		collection: db.GetCollection(companiesCollection),
	}
}

// CreateCompany 새 회사를 추가
func (r *CompanyRepository) CreateCompany(company entity.CompanyEntity) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.collection.InsertOne(ctx, company)
	return err
}

// GetAllCompanies 모든 회사 목록을 반환
func (r *CompanyRepository) GetAllCompanies() ([]entity.CompanyEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var companies []entity.CompanyEntity
	if err = cursor.All(ctx, &companies); err != nil {
		return nil, err
	}

	// nil 대신 빈 슬라이스 반환
	if companies == nil {
		companies = []entity.CompanyEntity{}
	}

	return companies, nil
}

// FindCompanyByUserID 사용자 ID로 회사를 찾음
func (r *CompanyRepository) FindCompanyByUserID(userID string) (*entity.CompanyEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var company entity.CompanyEntity
	filter := bson.D{{Key: "userId", Value: userID}}

	err := r.collection.FindOne(ctx, filter).Decode(&company)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // 문서가 없으면 nil 반환
		}
		return nil, err
	}

	return &company, nil
}

// FindCompanyByID ID로 회사를 찾음
func (r *CompanyRepository) FindCompanyByID(id string) (*entity.CompanyEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var company entity.CompanyEntity
	filter := bson.D{{Key: "id", Value: id}}

	err := r.collection.FindOne(ctx, filter).Decode(&company)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // 문서가 없으면 nil 반환
		}
		return nil, err
	}

	return &company, nil
}

// UpdateCompany 회사 정보를 업데이트
func (r *CompanyRepository) UpdateCompany(id string, updates bson.D) error {
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

// DeleteCompany 회사를 삭제
func (r *CompanyRepository) DeleteCompany(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{Key: "id", Value: id}}
	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}

// FindCompaniesByName 회사명으로 회사들을 찾음 (검색용)
func (r *CompanyRepository) FindCompaniesByName(companyName string) ([]entity.CompanyEntity, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 부분 일치 검색
	filter := bson.D{{Key: "companyName", Value: bson.D{{Key: "$regex", Value: companyName}, {Key: "$options", Value: "i"}}}}

	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var companies []entity.CompanyEntity
	if err = cursor.All(ctx, &companies); err != nil {
		return nil, err
	}

	if companies == nil {
		companies = []entity.CompanyEntity{}
	}

	return companies, nil
}
