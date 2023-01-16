package dto

import "github.com/issy20/go-simple-chat/domain/entity"

type UserDto struct {
	Id   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func UserDtoToEntity(dto *UserDto) *entity.User {
	return &entity.User{
		Id:   dto.Id,
		Name: dto.Name,
	}
}

func UserEntityToDto(u *entity.User) UserDto {
	return UserDto{
		Id:   u.Id,
		Name: u.Name,
	}
}
