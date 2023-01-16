package dto

import "github.com/issy20/go-simple-chat/domain/entity"

type MemberDto struct {
	RoomID int `db:"room_id"`
	UserID int `db:"user_id"`
}

func MemberDtoToEntity(dto *MemberDto) *entity.Member {
	return &entity.Member{
		RoomID: dto.RoomID,
		UserID: dto.UserID,
	}
}

func MemberEntityToDto(u *entity.Member) MemberDto {
	return MemberDto{
		RoomID: u.RoomID,
		UserID: u.UserID,
	}
}
