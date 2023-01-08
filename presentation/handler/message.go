package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/issy20/go-simple-chat/domain/entity"
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

func (mh *MessageHandler) MessagePost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	view := &JsonView{w: w, successCode: http.StatusCreated}

	w.Header().Add("Content-Type", "application/json")
	content, _ := ioutil.ReadAll(r.Body)

	var messageJson MessageJson
	if err := json.Unmarshal(content, &messageJson); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	message := messageJsonToEntity(&messageJson)
	fmt.Print("body")

	createdMessage, err := mh.mu.CreateMessage(ctx, message)

	if err != nil {
		logger.Println(err)
		view.ErrorView(err)
		return
	}
	res, err := json.Marshal(messageEntityToJson(createdMessage))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

type MessageJson struct {
	Id      int    `json:"id"`
	RoomID  int    `json:"room_id"`
	UserID  int    `json:"user_id"`
	Message string `json:"message"`
}

func messageEntityToJson(c *entity.Message) MessageJson {
	return MessageJson{
		Id:      c.Id,
		RoomID:  c.RoomID,
		UserID:  c.UserID,
		Message: c.Message,
	}
}

func messageJsonToEntity(j *MessageJson) *entity.Message {
	return &entity.Message{
		Id:      j.Id,
		RoomID:  j.RoomID,
		UserID:  j.UserID,
		Message: j.Message,
	}
}
