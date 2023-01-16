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
