package entity

type Message struct {
	Id      int
	RoomID  int
	UserID  int
	Message string
}

type Messages []*Message
