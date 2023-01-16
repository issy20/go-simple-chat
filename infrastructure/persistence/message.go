package persistence

import (
	"context"
	"fmt"

	"github.com/issy20/go-simple-chat/domain/entity"
	"github.com/issy20/go-simple-chat/domain/repository"
	"github.com/issy20/go-simple-chat/dto"
	"github.com/issy20/go-simple-chat/infrastructure/database"
)

var _ repository.IMessageRepository = &MessageRepository{}

type MessageRepository struct {
	conn *database.Conn
}

func NewMessageRepository(conn *database.Conn) repository.IMessageRepository {
	return &MessageRepository{
		conn: conn,
	}
}

func (mr *MessageRepository) CreateMessage(ctx context.Context, message *entity.Message) (*entity.Message, error) {
	query := `
		INSERT INTO messages (room_id, user_id, message)
		VALUES (?, ?, ?)
	`
	stmt, err := mr.conn.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	messageDto := dto.MessageEntityToDto(message)
	fmt.Println(&messageDto)
	res, err := stmt.ExecContext(ctx, &messageDto.RoomID, &messageDto.UserID, &messageDto.Message)

	id, _ := res.LastInsertId()
	messageDto.Id = (int)(id)

	if err != nil {
		return nil, fmt.Errorf("MessageRepository.CreateNewMessage ExecContext Error : %w", err)
	}

	return dto.MessageDtoToEntity(&messageDto), nil
}

func (mr *MessageRepository) GetMessagesByRoomID(ctx context.Context, roomId int) ([]*entity.Message, error) {
	var messageDto []*dto.MessageDto
	query := `
	  SELECT * FROM messages WHERE room_id = ? ORDER BY ID ASC
	`

	stmt, err := mr.conn.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	err = mr.conn.DB.SelectContext(ctx, &messageDto, query, roomId)
	if err != nil {
		return nil, fmt.Errorf("MessageRepository.GetMessagesByRoomID Get Error : %w", err)
	}
	return dto.MessagesDtoToEntity(messageDto), nil
}

func (mr *MessageRepository) GetRoomAndMessagesByRoomID(ctx context.Context, usersId string) ([]*entity.Message, error) {
	// get room id
	var roomMemberDto dto.RoomMemberDto
	query1 := `
		SELECT * FROM (
			SELECT room_id, GROUP_CONCAT(user_id ORDER BY user_id SEPARATOR ',') as room_member
			FROM members
			GROUP BY room_id
		) as member_room
		WHERE room_member = ?
  `
	stmt1, err := mr.conn.DB.Prepare(query1)
	if err != nil {
		return nil, err
	}
	defer stmt1.Close()

	err = mr.conn.DB.GetContext(ctx, &roomMemberDto, query1, usersId)
	if err != nil {
		return nil, fmt.Errorf("MessageRepository.GetRoomAndMessagesByRoomID Get roomId Error : %w", err)
	}

	// get message by room id
	var messageDto []*dto.MessageDto
	query2 := `
	  SELECT * FROM messages WHERE room_id = ? ORDER BY ID ASC
	`

	stmt2, err := mr.conn.DB.Prepare(query2)
	if err != nil {
		return nil, err
	}
	defer stmt2.Close()
	err = mr.conn.DB.SelectContext(ctx, &messageDto, query2, &roomMemberDto.RoomID)
	if err != nil {
		return nil, fmt.Errorf("MessageRepository.GetRoomAndMessagesByRoomID Get messages Error : %w", err)
	}

	return dto.MessagesDtoToEntity(messageDto), nil
}
