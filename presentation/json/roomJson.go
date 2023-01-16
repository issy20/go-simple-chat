package j

import "github.com/issy20/go-simple-chat/domain/entity"

type RoomJson struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func RoomsEntityToJson(e []*entity.Room) []RoomJson {
	jsonRooms := make([]RoomJson, len(e))
	for i, room := range e {
		jsonRooms[i] = RoomJson{
			Id:   room.Id,
			Name: room.Name,
		}
	}
	return jsonRooms
}

func RoomEntityToJson(c *entity.Room) RoomJson {
	return RoomJson{
		Id:   c.Id,
		Name: c.Name,
	}
}

type RoomMemberJson struct {
	RoomID     int    `json:"my_id"`
	RoomMember string `json:"room_member"`
}

func RoomMemberEntityToJson(j *entity.RoomMember) RoomMemberJson {
	return RoomMemberJson{
		RoomID:     j.RoomID,
		RoomMember: j.RoomMember,
	}
}

// func roomMemberJsonToEntity(e *RoomMemberJson) *entity.RoomMember {
// 	return &entity.RoomMember{
// 		RoomID:     e.RoomID,
// 		RoomMember: e.RoomMember,
// 	}
// }

type RoomMemberInputJson struct {
	MyID   int `json:"my_id"`
	UserID int `json:"user_id"`
}

func RoomMemberInputJsonToEntity(j *RoomMemberInputJson) *entity.GetRoomMemberInput {
	return &entity.GetRoomMemberInput{
		MyID:   j.MyID,
		UserID: j.UserID,
	}
}

// func roomMemberInputEntityToJson(e *entity.GetRoomMemberInput) RoomMemberInputJson {
// 	return RoomMemberInputJson{
// 		MyID:   e.MyID,
// 		UserID: e.UserID,
// 	}
// }
