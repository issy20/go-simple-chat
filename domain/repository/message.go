package repository

import (
	"context"

	"github.com/issy20/go-simple-chat/domain/entity"
)

type IMessageRepository interface {
	CreateMessage(ctx context.Context, message *entity.Message) (*entity.Message, error)
	GetMessagesByRoomID(ctx context.Context, roomId int) ([]*entity.Message, error)
}
