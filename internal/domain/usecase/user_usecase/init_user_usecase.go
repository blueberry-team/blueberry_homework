package user_usecase

import "blueberry_homework/internal/domain/repo_interface"

type UserUsecase struct {
	repo repo_interface.UserRepository
}

func NewUserUsecase(r repo_interface.UserRepository) *UserUsecase {
	return &UserUsecase{
		repo: r,
	}
}
