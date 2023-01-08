package usecase

import (
	"context"
	"fmt"

	"github.com/issy20/go-simple-chat/domain/entity"
	"github.com/issy20/go-simple-chat/domain/repository"
)

var _ IRoomUsecase = &RoomUsecase{}

type RoomUsecase struct {
	repo repository.IRoomRepository
}

type IRoomUsecase interface {
	CreateRoom(ctx context.Context, room *entity.Room) (*entity.Room, error)
	GetRoomByUsersID(ctx context.Context, input *entity.GetRoomMemberInput) (*entity.RoomMember, error)
}

func NewRoomUsecase(rr repository.IRoomRepository) IRoomUsecase {
	return &RoomUsecase{
		repo: rr,
	}
}

func (rr *RoomUsecase) CreateRoom(ctx context.Context, room *entity.Room) (*entity.Room, error) {
	return rr.repo.CreateRoom(ctx, room)
}

func (rr *RoomUsecase) GetRoomByUsersID(ctx context.Context, input *entity.GetRoomMemberInput) (*entity.RoomMember, error) {
	usersId := fmt.Sprintf("%d,%d", input.MyID, input.UserID)
	fmt.Print(usersId)
	return rr.repo.GetRoomByUsersID(ctx, usersId)
}
