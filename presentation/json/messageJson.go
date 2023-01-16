package j

import "github.com/issy20/go-simple-chat/domain/entity"

type MessageJson struct {
	Id      int    `json:"id"`
	RoomID  int    `json:"room_id"`
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}

func MessageEntityToJson(c *entity.Message) MessageJson {
	return MessageJson{
		Id:      c.Id,
		RoomID:  c.RoomID,
		UserID:  c.UserID,
		Message: c.Message,
	}
}

func MessageJsonToEntity(j *MessageJson) *entity.Message {
	return &entity.Message{
		Id:      j.Id,
		RoomID:  j.RoomID,
		UserID:  j.UserID,
		Message: j.Message,
	}
}

func MessagesEntityToJson(e []*entity.Message) []MessageJson {
	messagesJson := make([]MessageJson, len(e))
	for i, message := range e {
		messagesJson[i] = MessageJson{
			Id:      message.Id,
			RoomID:  message.RoomID,
			UserID:  message.UserID,
			Message: message.Message,
		}
	}
	return messagesJson
}
