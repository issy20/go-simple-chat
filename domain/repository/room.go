package repository

import (
	"context"

	"github.com/issy20/go-simple-chat/domain/entity"
)

type IRoomRepository interface {
	CreateRoom(ctx context.Context, room *entity.Room) (*entity.Room, error)
	GetRoomByUsersID(ctx context.Context, usersID string) (*entity.RoomMember, error)
}
