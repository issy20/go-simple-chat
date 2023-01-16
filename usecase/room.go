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
	GetAllRoomNameByUserID(ctx context.Context, userID int) ([]*entity.Room, error)
}

func NewRoomUsecase(rr repository.IRoomRepository) IRoomUsecase {
	return &RoomUsecase{
		repo: rr,
	}
}

func (ru *RoomUsecase) CreateRoom(ctx context.Context, room *entity.Room) (*entity.Room, error) {
	return ru.repo.CreateRoom(ctx, room)
}

func (ru *RoomUsecase) GetRoomByUsersID(ctx context.Context, input *entity.GetRoomMemberInput) (*entity.RoomMember, error) {
	usersId := fmt.Sprintf("%d,%d", input.MyID, input.UserID)
	fmt.Print(usersId)
	return ru.repo.GetRoomByUsersID(ctx, usersId)
}

func (ru *RoomUsecase) GetAllRoomNameByUserID(ctx context.Context, userID int) ([]*entity.Room, error) {
	return ru.repo.GetAllRoomNameByUserID(ctx, userID)
}
