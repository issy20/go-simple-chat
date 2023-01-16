package dto

import "github.com/issy20/go-simple-chat/domain/entity"

type RoomDto struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

func RoomDtoToEntity(dto *RoomDto) *entity.Room {
	return &entity.Room{
		Id:   dto.Id,
		Name: dto.Name,
	}
}

func RoomsDtoToEntity(dto []*RoomDto) []*entity.Room {
	entityRoom := make([]*entity.Room, len(dto))
	for i, room := range dto {
		entityRoom[i] = &entity.Room{
			Id:   room.Id,
			Name: room.Name,
		}
	}
	return entityRoom
}

func RoomEntityToDto(r *entity.Room) RoomDto {
	return RoomDto{
		Id:   r.Id,
		Name: r.Name,
	}
}

type RoomMemberDto struct {
	RoomID     int    `db:"room_id"`
	RoomMember string `db:"room_member"`
}

func RoomMemberDtoToEntity(dto *RoomMemberDto) *entity.RoomMember {
	return &entity.RoomMember{
		RoomID:     dto.RoomID,
		RoomMember: dto.RoomMember,
	}
}

// func roomMemberEntityToDto(r *entity.RoomMember) roomMemberDto {
// 	return roomMemberDto{
// 		RoomID:     r.RoomID,
// 		RoomMember: r.RoomMember,
// 	}
// }
