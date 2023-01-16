package persistence

import (
	"context"
	"fmt"

	"github.com/issy20/go-simple-chat/domain/entity"
	"github.com/issy20/go-simple-chat/domain/repository"
	"github.com/issy20/go-simple-chat/dto"
	"github.com/issy20/go-simple-chat/infrastructure/database"
)

var _ repository.IMemberRepository = &MemberRepository{}

type MemberRepository struct {
	conn *database.Conn
}

func NewMemberRepository(conn *database.Conn) repository.IMemberRepository {
	return &MemberRepository{
		conn: conn,
	}
}

func (ur *MemberRepository) CreateMember(ctx context.Context, member *entity.Member) (*entity.Member, error) {
	query := `
		INSERT INTO members (room_id, user_id)
		VALUES (:room_id, :user_id);
	`
	// stmt, err := ur.conn.DB.Prepare(query)
	// if err != nil {
	// 	return nil, err
	// }
	// defer stmt.Close()
	// dto := memberEntityToDto(member)
	// if _, err := stmt.ExecContext(ctx, &dto.RoomID, &dto.UserID); err != nil {
	// 	return nil, err
	// }

	// if err != nil {
	// 	return nil, fmt.Errorf("MemberRepository.CreateMember ExecContext Error : %w", err)
	// }
	memberDto := dto.MemberEntityToDto(member)

	// fmt.Print(&dto)

	// params := map[string]interface{}{
	// 	"room_id": &dto.RoomID,
	// 	"user_id": &dto.UserID,
	// }

	if _, err := ur.conn.DB.NamedExecContext(ctx, query, &memberDto); err != nil {
		return nil, fmt.Errorf("MemberRepository.CreateMember ExecContext Error : %w", err)
	}

	return dto.MemberDtoToEntity(&memberDto), nil
}
