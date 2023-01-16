package usecase

import (
	"context"
	"fmt"

	"github.com/issy20/go-simple-chat/domain/entity"
	"github.com/issy20/go-simple-chat/domain/repository"
)

var _ IMessageUsecase = &MessageUsecase{}

type MessageUsecase struct {
	repo repository.IMessageRepository
}

type IMessageUsecase interface {
	CreateMessage(ctx context.Context, message *entity.Message) (*entity.Message, error)
	GetMessagesByRoomID(ctx context.Context, roomId int) ([]*entity.Message, error)
}

func NewMessageUsecase(mr repository.IMessageRepository) IMessageUsecase {
	return &MessageUsecase{
		repo: mr,
	}
}

func (mu *MessageUsecase) CreateMessage(ctx context.Context, message *entity.Message) (*entity.Message, error) {
	fmt.Print(message)
	return mu.repo.CreateMessage(ctx, message)
}

func (mu *MessageUsecase) GetMessagesByRoomID(ctx context.Context, roomId int) ([]*entity.Message, error) {
	return mu.repo.GetMessagesByRoomID(ctx, roomId)
}
