package persistence

import (
	"context"
	"fmt"

	"github.com/issy20/go-simple-chat/domain/entity"
	"github.com/issy20/go-simple-chat/domain/repository"
	"github.com/issy20/go-simple-chat/dto"
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
		VALUES (?)
	`
	stmt, err := rr.conn.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	roomDto := dto.RoomEntityToDto(room)
	res, err := stmt.ExecContext(ctx, &roomDto.Name)

	id, _ := res.LastInsertId()
	roomDto.Id = (int)(id)

	if err != nil {
		return nil, fmt.Errorf("RoomRepository.CreateNewRoom ExecContext Error : %w", err)
	}

	return dto.RoomDtoToEntity(&roomDto), nil
}

func (rr *RoomRepository) GetRoomByUsersID(ctx context.Context, usersId string) (*entity.RoomMember, error) {
	var roomMemberDto dto.RoomMemberDto
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

	err = rr.conn.DB.GetContext(ctx, &roomMemberDto, query, usersId)
	if err != nil {
		return nil, fmt.Errorf("RoomRepository.GetRoomByUsersID Get Error : %w", err)
	}

	return dto.RoomMemberDtoToEntity(&roomMemberDto), nil
}

func (rr *RoomRepository) GetAllRoomNameByUserID(ctx context.Context, userID int) ([]*entity.Room, error) {
	var roomDto []*dto.RoomDto
	query := `
	  SELECT members.room_id as id, rooms.name as name FROM members
		INNER JOIN rooms ON rooms.id = members.room_id
		WHERE user_id = ?
	`
	stmt, err := rr.conn.DB.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	err = rr.conn.DB.SelectContext(ctx, &roomDto, query, userID)
	fmt.Print("room_member")
	if err != nil {
		return nil, fmt.Errorf("RoomRepository.GetAllRoomNameByUserID Get Error : %w", err)
	}
	return dto.RoomsDtoToEntity(roomDto), nil
}
