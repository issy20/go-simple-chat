package repository

import (
	"context"

	"github.com/issy20/go-simple-chat/domain/entity"
)

type IUserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
}
