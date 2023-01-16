package j

import "github.com/issy20/go-simple-chat/domain/entity"

type MemberJson struct {
	RoomID int `json:"room_id"`
	UserID int `json:"user_id"`
}

func MemberEntityToJson(c *entity.Member) MemberJson {
	return MemberJson{
		RoomID: c.RoomID,
		UserID: c.UserID,
	}
}

func MemberJsonToEntity(j *MemberJson) *entity.Member {
	return &entity.Member{
		RoomID: j.RoomID,
		UserID: j.UserID,
	}
}
