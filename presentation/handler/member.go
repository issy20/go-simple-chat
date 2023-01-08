package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/issy20/go-simple-chat/domain/entity"
	"github.com/issy20/go-simple-chat/usecase"
)

type MemberHandler struct {
	mu usecase.IMemberUsecase
}

func NewMemberHandler(mu usecase.IMemberUsecase) *MemberHandler {
	return &MemberHandler{
		mu: mu,
	}
}

func (mh *MemberHandler) MemberPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	view := &JsonView{w: w, successCode: http.StatusCreated}
	w.Header().Add("Content-Type", "application/json")
	content, _ := ioutil.ReadAll(r.Body)
	// json
	var memberJson MemberJson
	if err := json.Unmarshal(content, &memberJson); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// json -> entity
	member := memberJsonToEntity(&memberJson)
	createdMember, err := mh.mu.CreateMember(ctx, member)
	if err != nil {
		logger.Println(err)
		view.ErrorView(err)
		return
	}
	res, err := json.Marshal(memberEntityToJson(createdMember))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}

type MemberJson struct {
	RoomID int `json:"room_id"`
	UserID int `json:"user_id"`
}

func memberEntityToJson(c *entity.Member) MemberJson {
	return MemberJson{
		RoomID: c.RoomID,
		UserID: c.UserID,
	}
}

func memberJsonToEntity(j *MemberJson) *entity.Member {
	return &entity.Member{
		RoomID: j.RoomID,
		UserID: j.UserID,
	}
}
