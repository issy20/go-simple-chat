package persistence

import (
	"context"
	"fmt"

	"github.com/issy20/go-simple-chat/domain/entity"
	"github.com/issy20/go-simple-chat/domain/repository"
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
	dto := messageEntityToDto(message)
	fmt.Println(&dto)
	res, err := stmt.ExecContext(ctx, &dto.RoomID, &dto.UserID, &dto.Message)

	id, _ := res.LastInsertId()
	dto.Id = (int)(id)

	if err != nil {
		return nil, fmt.Errorf("MessageRepository.CreateNewMessage ExecContext Error : %w", err)
	}

	return messageDtoToEntity(&dto), nil
}

type messageDto struct {
	Id      int    `db:"id" json:"id"`
	RoomID  int    `db:"room_id" json:"room_id"`
	UserID  int    `db:"user_id" json:"user_id"`
	Message string `db:"message" json:"message"`
}

func messageDtoToEntity(dto *messageDto) *entity.Message {
	return &entity.Message{
		Id:      dto.Id,
		RoomID:  dto.RoomID,
		UserID:  dto.UserID,
		Message: dto.Message,
	}
}

func messageEntityToDto(u *entity.Message) messageDto {
	return messageDto{
		Id:      u.Id,
		RoomID:  u.RoomID,
		UserID:  u.UserID,
		Message: u.Message,
	}
}
