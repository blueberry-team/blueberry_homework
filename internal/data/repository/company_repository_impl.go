package repository

import (
	"blueberry_homework/dto/response"
	"blueberry_homework/internal/domain/entities"
	"blueberry_homework/internal/domain/repo_interface"
	"fmt"

	"github.com/gocql/gocql"
)

// companyRepo CompanyRepository 인터페이스의 구현체입니다.
type companyRepo struct {
	// 저장소
	session *gocql.Session
}

// NewCompanyRepository 새로운 CompanyRepository 인스턴스를 생성합니다.
func NewCompanyRepository(session *gocql.Session) repo_interface.CompanyRepository {
	return &companyRepo{
		session: session,
	}
}

// CheckCompanyWithUserId: userId로 회사 존재 여부 확인
func (r *companyRepo) CheckCompanyWithUserId(userId gocql.UUID) (bool, error) {
	var dummy gocql.UUID
	err := r.session.Query(`
		SELECT id FROM companies WHERE user_id = ? LIMIT 1
	`, userId).Scan(&dummy)
	if err != nil {
		if err == gocql.ErrNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// Company entity 를 저장소에 추가하는 함수
func (r *companyRepo) CreateCompany(entity entities.CompanyEntity) error {
	return r.session.Query(`
		INSERT INTO companies (id, user_id, company_name, company_address, total_staff, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?)
	`, entity.Id, entity.UserId, entity.CompanyName, entity.CompanyAddress, entity.TotalStaff, entity.CreatedAt, entity.UpdatedAt).Exec()
}

// GetUserCompany: userId로 회사 정보 반환
func (r *companyRepo) GetUserCompany(userId gocql.UUID) (response.CompanyResponse, error) {
	var entity entities.CompanyEntity
	err := r.session.Query(`
		SELECT id, user_id, company_name, company_address, total_staff, created_at, updated_at FROM companies WHERE user_id = ? LIMIT 1
	`, userId).Scan(
		&entity.Id, &entity.UserId, &entity.CompanyName, &entity.CompanyAddress, &entity.TotalStaff, &entity.CreatedAt, &entity.UpdatedAt,
	)
	if err != nil {
		if err == gocql.ErrNotFound {
			return response.CompanyResponse{}, fmt.Errorf("company not found for userId: %s", userId.String())
		}
		return response.CompanyResponse{}, err
	}

	return response.CompanyResponse{
		Id:             entity.Id,
		UserId:         entity.UserId,
		CompanyName:    entity.CompanyName,
		CompanyAddress: entity.CompanyAddress,
		TotalStaff:     entity.TotalStaff,
		CreatedAt:      entity.CreatedAt,
		UpdatedAt:      entity.UpdatedAt,
	}, nil
}

// getCompanyId: userId로 회사 id 반환
// 당장은 쓸 일이 없어서 private 으로 생성되어있음
func (r *companyRepo) getCompanyId(userId gocql.UUID) (gocql.UUID, error) {
	var companyId gocql.UUID
	err := r.session.Query(`
		SELECT id FROM companies WHERE user_id = ? LIMIT 1
	`, userId).Scan(&companyId)
	return companyId, err
}

// UpdateCompany는 회사 정보를 수정합니다.
func (r *companyRepo) ChangeCompany(entity entities.ChangeCompanyEntity) error {
	var companyId gocql.UUID

	companyId, err := r.getCompanyId(entity.UserId)
	if err != nil {
		return err
	}

	return r.session.Query(`
		UPDATE companies
		SET company_name = ?, company_address = ?, total_staff = ?, updated_at = ?
		WHERE id = ?
	`, entity.CompanyName, entity.CompanyAddress, entity.TotalStaff, entity.UpdatedAt, companyId).Exec()
}

// DeleteCompany: userId로 회사 삭제
func (r *companyRepo) DeleteCompany(userId gocql.UUID) error {
	var companyId gocql.UUID

	companyId, err := r.getCompanyId(userId)
	if err != nil {
		return err
	}

	return r.session.Query(`
		DELETE FROM companies WHERE id = ?
	`, companyId).Exec()
}
