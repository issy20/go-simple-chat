package persistence

import (
	"context"
	"fmt"

	"github.com/issy20/go-simple-chat/domain/entity"
	"github.com/issy20/go-simple-chat/domain/repository"
	"github.com/issy20/go-simple-chat/infrastructure/database"
)

var _ repository.IRoomRepository = &RoomRepository{}

type RoomRepository struct {
	conn *database.Conn
}

func NewRoomRepository(conn *database.Conn) repository.IRoomRepository {
	return &RoomRepository{
		conn: conn,
	}
}

func (rr *RoomRepository) CreateRoom(ctx context.Context, room *entity.Room) (*entity.Room, error) {
	query := `
	INSERT INTO rooms (name)
		VALUES (:name)
	`
	stmt, err := rr.conn.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	dto := roomEntityToDto(room)
	res, err := stmt.ExecContext(ctx, &dto.Name)

	id, _ := res.LastInsertId()
	dto.Id = (int)(id)

	if err != nil {
		return nil, fmt.Errorf("RoomRepository.CreateNewRoom ExecContext Error : %w", err)
	}

	return roomDtoToEntity(&dto), nil

}

func (rr *RoomRepository) GetRoomByUsersID(ctx context.Context, usersId string) (*entity.RoomMember, error) {
	var dto roomMemberDto
	query := `
		SELECT * FROM (
			SELECT room_id, GROUP_CONCAT(user_id ORDER BY user_id SEPARATOR ',') as room_member
			FROM members
			GROUP BY room_id
		) as member_room
		WHERE room_member = ?
	`

	stmt, err := rr.conn.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = rr.conn.DB.GetContext(ctx, &dto, query, usersId)
	if err != nil {
		return nil, fmt.Errorf("RoomRepository.GetRoomByUsersID Get Error : %w", err)
	}

	return roomMemberDtoToEntity(&dto), nil

}

type roomDto struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

func roomDtoToEntity(dto *roomDto) *entity.Room {
	return &entity.Room{
		Id:   dto.Id,
		Name: dto.Name,
	}
}

func roomEntityToDto(r *entity.Room) roomDto {
	return roomDto{
		Id:   r.Id,
		Name: r.Name,
	}
}

type roomMemberDto struct {
	RoomID     int    `db:"room_id"`
	RoomMember string `db:"room_member"`
}

func roomMemberDtoToEntity(dto *roomMemberDto) *entity.RoomMember {
	return &entity.RoomMember{
		RoomID:     dto.RoomID,
		RoomMember: dto.RoomMember,
	}
}

// func roomMemberEntityToDto(r *entity.RoomMember) roomMemberDto {
// 	return roomMemberDto{
// 		RoomID:     r.RoomID,
// 		RoomMember: r.RoomMember,
// 	}
// }
