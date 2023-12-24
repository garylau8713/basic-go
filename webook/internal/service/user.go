package service

import (
	"basic-go/webook/internal/domain"
	"basic-go/webook/internal/repository"
	"context"
)

type UserService struct {
	// TODO: The reason that we choose to use pointer is
	// we are trying to keep only one repo instance for this whole project
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// In the Service Directory, the method name should keep align with the biz logic
// So keep it as SignUp instead of something like Creat/Insert.

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	return svc.repo.Create(ctx, u)
}
