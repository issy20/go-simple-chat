package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog"
	"github.com/issy20/go-simple-chat/infrastructure/database"
	"github.com/issy20/go-simple-chat/infrastructure/persistence"
	"github.com/issy20/go-simple-chat/presentation/handler"
	"github.com/issy20/go-simple-chat/usecase"
)

func main() {
	conn, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := conn.DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	l := httplog.NewLogger("app", httplog.Options{
		JSON: true,
	})

	ur := persistence.NewUserRepository(conn)
	uc := usecase.NewUserUsecase(ur)
	uh := handler.NewUserHandler(uc)

	rr := persistence.NewRoomRepository(conn)
	rc := usecase.NewRoomUsecase(rr)
	rh := handler.NewRoomHandler(rc)

	mr := persistence.NewMessageRepository(conn)
	mc := usecase.NewMessageUsecase(mr)
	mh := handler.NewMessageHandler(mc)

	mmr := persistence.NewMemberRepository(conn)
	mmc := usecase.NewMemberUsecase(mmr)
	mmh := handler.NewMemberHandler(mmc)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	// r.Use(middleware.Logger)
	r.Use(httplog.RequestLogger(l))

	r.Route("/test", func(r chi.Router) {
		r.Post("/", handler.Test)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", uh.UserPost)
	})

	r.Route("/rooms", func(r chi.Router) {
		r.Post("/", rh.RoomPost)
		r.Get("/", rh.RoomGet)
	})

	r.Route("/messages", func(r chi.Router) {
		r.Post("/", mh.MessagePost)
	})

	r.Route("/members", func(r chi.Router) {
		r.Post("/", mmh.MemberPost)
	})

	http.ListenAndServe(":8080", r)

}
