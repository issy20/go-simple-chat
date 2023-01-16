package j

import "github.com/issy20/go-simple-chat/domain/entity"

type UserJson struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func UserEntityToJson(c *entity.User) UserJson {
	return UserJson{
		Id:   c.Id,
		Name: c.Name,
	}
}
