package repository

import (
	"basic-go/webook/internal/domain"
	"context"
)

type UserRepository struct {
}

func (repo *UserRepository) Create(ctx context.Context, u domain.User) error {

}
