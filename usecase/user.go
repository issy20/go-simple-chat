package usecase

import (
	"context"

	"github.com/issy20/go-simple-chat/domain/entity"
	"github.com/issy20/go-simple-chat/domain/repository"
)

var _ IUserUsecase = &UserUsecase{}

type UserUsecase struct {
	repo repository.IUserRepository
}

type IUserUsecase interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
}

func NewUserUsecase(ur repository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		repo: ur,
	}
}

func (uu *UserUsecase) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	return uu.repo.CreateUser(ctx, user)
}
