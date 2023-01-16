package dto

import "github.com/issy20/go-simple-chat/domain/entity"

type MessageDto struct {
	Id      int    `db:"id" json:"id"`
	RoomID  int    `db:"room_id" json:"room_id"`
	UserID  int    `db:"user_id" json:"user_id"`
	Message string `db:"message" json:"message"`
}

func MessageDtoToEntity(dto *MessageDto) *entity.Message {
	return &entity.Message{
		Id:      dto.Id,
		RoomID:  dto.RoomID,
		UserID:  dto.UserID,
		Message: dto.Message,
	}
}

func MessageEntityToDto(u *entity.Message) MessageDto {
	return MessageDto{
		Id:      u.Id,
		RoomID:  u.RoomID,
		UserID:  u.UserID,
		Message: u.Message,
	}
}

func MessagesDtoToEntity(dto []*MessageDto) []*entity.Message {
	MessageEntity := make([]*entity.Message, len(dto))
	for i, message := range dto {
		MessageEntity[i] = &entity.Message{
			Id:      message.Id,
			RoomID:  message.RoomID,
			UserID:  message.UserID,
			Message: message.Message,
		}
	}
	return MessageEntity
}
