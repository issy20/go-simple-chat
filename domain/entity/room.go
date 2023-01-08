package entity

type Room struct {
	Id   int
	Name string
}

type RoomMember struct {
	RoomID     int
	RoomMember string
}

type GetRoomMemberInput struct {
	MyID   int
	UserID int
}

type Rooms []*Room
