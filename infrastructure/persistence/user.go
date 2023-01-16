package persistence

import (
	"context"
	"fmt"

	"github.com/issy20/go-simple-chat/domain/entity"
	"github.com/issy20/go-simple-chat/domain/repository"
	"github.com/issy20/go-simple-chat/dto"
	"github.com/issy20/go-simple-chat/infrastructure/database"
)

var _ repository.IUserRepository = &UserRepository{}

type UserRepository struct {
	conn *database.Conn
}

func NewUserRepository(conn *database.Conn) repository.IUserRepository {
	return &UserRepository{
		conn: conn,
	}
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	query := `
		INSERT INTO users (name)
		VALUES (?)
	`
	stmt, err := ur.conn.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	userDto := dto.UserEntityToDto(user)
	res, err := stmt.ExecContext(ctx, &userDto.Name)

	id, _ := res.LastInsertId()
	userDto.Id = (int)(id)

	if err != nil {
		return nil, fmt.Errorf("UserRepository.CreateNewUser ExecContext Error : %w", err)
	}

	return dto.UserDtoToEntity(&userDto), nil
}
