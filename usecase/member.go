package usecase

import (
	"context"
	"fmt"

	"github.com/issy20/go-simple-chat/domain/entity"
	"github.com/issy20/go-simple-chat/domain/repository"
)

var _ IMemberUsecase = &MemberUsecase{}

type MemberUsecase struct {
	repo repository.IMemberRepository
}

type IMemberUsecase interface {
	CreateMember(ctx context.Context, member *entity.Member) (*entity.Member, error)
}

func NewMemberUsecase(mr repository.IMemberRepository) IMemberUsecase {
	return &MemberUsecase{
		repo: mr,
	}
}

func (mu *MemberUsecase) CreateMember(ctx context.Context, member *entity.Member) (*entity.Member, error) {
	fmt.Print(member)

	return mu.repo.CreateMember(ctx, member)
}
