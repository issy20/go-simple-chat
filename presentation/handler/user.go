package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/issy20/go-simple-chat/domain/entity"
	j "github.com/issy20/go-simple-chat/presentation/json"
	"github.com/issy20/go-simple-chat/usecase"
)

type UserHandler struct {
	uu usecase.IUserUsecase
}

func NewUserHandler(uu usecase.IUserUsecase) *UserHandler {
	return &UserHandler{
		uu: uu,
	}
}

func (u *UserHandler) UserPost(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	view := &JsonView{w: w, successCode: http.StatusCreated}
	user := &entity.User{}
	fmt.Println(json.NewDecoder(r.Body))
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		logger.Println(err)
		view.ErrorView(err)
		return
	}
	user, err := u.uu.CreateUser(ctx, user)
	if err != nil {
		logger.Println(err)
		view.ErrorView(err)
		return
	}
	view.View(j.UserEntityToJson(user))
}
