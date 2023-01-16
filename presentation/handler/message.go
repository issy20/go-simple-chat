package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	j "github.com/issy20/go-simple-chat/presentation/json"
	"github.com/issy20/go-simple-chat/usecase"
)

type MessageHandler struct {
	mu usecase.IMessageUsecase
}

func NewMessageHandler(mu usecase.IMessageUsecase) *MessageHandler {
	return &MessageHandler{
		mu: mu,
	}
}

func (mh *MessageHandler) PostMessage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	view := &JsonView{w: w, successCode: http.StatusCreated}

	w.Header().Add("Content-Type", "application/json")
	content, _ := ioutil.ReadAll(r.Body)

	var messageJson j.MessageJson
	if err := json.Unmarshal(content, &messageJson); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	message := j.MessageJsonToEntity(&messageJson)
	fmt.Print("body")

	createdMessage, err := mh.mu.CreateMessage(ctx, message)

	if err != nil {
		logger.Println(err)
		view.ErrorView(err)
		return
	}
	res, err := json.Marshal(j.MessageEntityToJson(createdMessage))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func (mh *MessageHandler) GetMessagesByRoomID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	view := &JsonView{w: w, successCode: http.StatusCreated}
	w.Header().Add("Content-Type", "application/json")
	roomId, _ := strconv.Atoi(chi.URLParam(r, "room_id"))

	fmt.Print("roomId", roomId)
	messages, err := mh.mu.GetMessagesByRoomID(ctx, roomId)
	if err != nil {
		logger.Println(err)
		view.ErrorView(err)
		return
	}
	res, err := json.Marshal(j.MessagesEntityToJson(messages))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

func (mh *MessageHandler) GetRoomAndMessagesByRoomID(w http.ResponseWriter, r *http.Request) {
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
	messages, err := mh.mu.GetRoomAndMessagesByRoomID(ctx, roomMemberInput)
	fmt.Print("roomMember", messages)

	if err != nil {
		logger.Println(err)
		view.ErrorView(err)
		return
	}
	res, err := json.Marshal(j.MessagesEntityToJson(messages))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}
