package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/issy20/go-simple-chat/domain/entity"
	j "github.com/issy20/go-simple-chat/presentation/json"

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

func (rh *RoomHandler) GetRoomMember(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	view := &JsonView{w: w, successCode: http.StatusCreated}
	w.Header().Add("Content-Type", "application/json")
	content, _ := ioutil.ReadAll(r.Body)
	var roomMemberInputJson j.RoomMemberInputJson
	if err := json.Unmarshal(content, &roomMemberInputJson); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	roomMemberInput := j.RoomMemberInputJsonToEntity(&roomMemberInputJson)
	fmt.Print("roomMemberInput", roomMemberInput)
	roomMember, err := rh.ru.GetRoomByUsersID(ctx, roomMemberInput)
	fmt.Print("roomMember", roomMember)

	if err != nil {
		logger.Println(err)
		view.ErrorView(err)
		return
	}
	res, err := json.Marshal(j.RoomMemberEntityToJson(roomMember))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func (rh *RoomHandler) GetAllRoomName(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	view := &JsonView{w: w, successCode: http.StatusCreated}
	w.Header().Add("Content-Type", "application/json")
	userId, _ := strconv.Atoi(chi.URLParam(r, "user_id"))

	fmt.Print("userid", userId)
	rooms, err := rh.ru.GetAllRoomNameByUserID(ctx, userId)
	if err != nil {
		logger.Println(err)
		view.ErrorView(err)
		return
	}
	res, err := json.Marshal(j.RoomsEntityToJson(rooms))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func (rh *RoomHandler) PostRoom(w http.ResponseWriter, r *http.Request) {
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
	view.View(j.RoomEntityToJson(room))
}
