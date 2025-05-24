package usecase

import (
	"errors"

	"blueberry_homework_go_gin/entity"
	"blueberry_homework_go_gin/repository"
	"blueberry_homework_go_gin/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// AuthUseCase 인증 관련 비즈니스 로직을 담당하는 유스케이스
type AuthUseCase struct {
	authRepo *repository.AuthRepository
}

// NewAuthUseCase 새로운 AuthUseCase 인스턴스를 생성
func NewAuthUseCase(authRepo *repository.AuthRepository) *AuthUseCase {
	return &AuthUseCase{
		authRepo: authRepo,
	}
}

// SignUp 회원가입
func (u *AuthUseCase) SignUp(email, password, name, role string) (*entity.UserPublicInfo, error) {
	// 이메일 형식 검증
	if !utils.ValidateEmail(email) {
		return nil, errors.New("invalid email format")
	}

	// 비밀번호 강도 검증
	if err := utils.ValidatePassword(password); err != nil {
		return nil, err
	}

	// 역할 검증
	if !entity.IsValidRole(role) {
		return nil, errors.New("invalid role. must be 'boss' or 'worker'")
	}

	// 이메일 중복 확인
	existingUser, err := u.authRepo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already exists")
	}

	// 비밀번호 해싱
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// 새 사용자 생성
	user := entity.NewUserEntity(email, hashedPassword, name, entity.UserRole(role))
	if err := u.authRepo.CreateUser(user); err != nil {
		return nil, err
	}

	// 공개 정보 반환
	publicInfo := user.ToPublicInfo()
	return &publicInfo, nil
}

// LogIn 로그인
func (u *AuthUseCase) LogIn(email, password string) (*entity.UserPublicInfo, error) {
	// 사용자 찾기
	user, err := u.authRepo.FindUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid email or password")
	}

	// 비밀번호 검증
	if !utils.VerifyPassword(password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	// 공개 정보 반환
	publicInfo := user.ToPublicInfo()
	return &publicInfo, nil
}

// GetUser 사용자 정보 조회
func (u *AuthUseCase) GetUser(userID string) (*entity.UserPublicInfo, error) {
	user, err := u.authRepo.FindUserByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	publicInfo := user.ToPublicInfo()
	return &publicInfo, nil
}

// ChangeUser 사용자 정보 수정
func (u *AuthUseCase) ChangeUser(userID, name, role string) (*entity.UserPublicInfo, error) {
	// 사용자 존재 확인
	user, err := u.authRepo.FindUserByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	// 업데이트할 필드 준비
	updates := bson.D{}

	if name != "" {
		updates = append(updates, bson.E{Key: "name", Value: name})
	}

	if role != "" {
		if !entity.IsValidRole(role) {
			return nil, errors.New("invalid role. must be 'boss' or 'worker'")
		}
		updates = append(updates, bson.E{Key: "role", Value: role})
	}

	// 업데이트 실행
	if err := u.authRepo.UpdateUser(userID, updates); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// 업데이트된 사용자 정보 조회
	return u.GetUser(userID)
}

// FindUserByID ID로 사용자 찾기 (내부 사용)
func (u *AuthUseCase) FindUserByID(userID string) (*entity.UserEntity, error) {
	return u.authRepo.FindUserByID(userID)
}
