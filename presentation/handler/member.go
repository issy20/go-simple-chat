package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	j "github.com/issy20/go-simple-chat/presentation/json"
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
	var memberJson j.MemberJson
	if err := json.Unmarshal(content, &memberJson); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// json -> entity
	member := j.MemberJsonToEntity(&memberJson)
	createdMember, err := mh.mu.CreateMember(ctx, member)
	if err != nil {
		logger.Println(err)
		view.ErrorView(err)
		return
	}
	res, err := json.Marshal(j.MemberEntityToJson(createdMember))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(res)
}
