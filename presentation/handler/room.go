package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/issy20/go-simple-chat/domain/entity"
	"github.com/issy20/go-simple-chat/usecase"
)

type RoomHandler struct {
	ru usecase.IRoomUsecase
}

func NewRoomHandler(ru usecase.IRoomUsecase) *RoomHandler {
	return &RoomHandler{
		ru,
	}
}

func (rh *RoomHandler) RoomGet(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	view := &JsonView{w: w, successCode: http.StatusCreated}
	w.Header().Add("Content-Type", "application/json")
	content, _ := ioutil.ReadAll(r.Body)
	var roomMemberInputJson RoomMemberInputJson
	if err := json.Unmarshal(content, &roomMemberInputJson); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	roomMemberInput := roomMemberInputJsonToEntity(&roomMemberInputJson)
	fmt.Print("roomMemberInput", roomMemberInput)
	roomMember, err := rh.ru.GetRoomByUsersID(ctx, roomMemberInput)
	fmt.Print("roomMember", roomMember)

	if err != nil {
		logger.Println(err)
		view.ErrorView(err)
		return
	}
	res, err := json.Marshal(roomMemberEntityToJson(roomMember))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func (rh *RoomHandler) RoomPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	view := &JsonView{w: w, successCode: http.StatusCreated}
	room := &entity.Room{}
	if err := json.NewDecoder(r.Body).Decode(room); err != nil {
		logger.Println(err)
		view.ErrorView(err)
		return
	}
	room, err := rh.ru.CreateRoom(ctx, room)
	if err != nil {
		logger.Println(err)
		view.ErrorView(err)
		return
	}
	view.View(roomEntityToJson(room))
}

type RoomJson struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func roomEntityToJson(c *entity.Room) RoomJson {
	return RoomJson{
		Id:   c.Id,
		Name: c.Name,
	}
}

type RoomMemberJson struct {
	RoomID     int    `json:"my_id"`
	RoomMember string `json:"room_member"`
}

func roomMemberEntityToJson(j *entity.RoomMember) RoomMemberJson {
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

func roomMemberInputJsonToEntity(j *RoomMemberInputJson) *entity.GetRoomMemberInput {
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
