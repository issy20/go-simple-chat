package repository

import (
	"context"

	"github.com/issy20/go-simple-chat/domain/entity"
)

type IMemberRepository interface {
	CreateMember(ctx context.Context, member *entity.Member) (*entity.Member, error)
}
