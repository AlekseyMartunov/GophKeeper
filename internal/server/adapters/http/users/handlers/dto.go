package userhandlers

import "GophKeeper/internal/server/entity/users"

type userDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (dto *userDTO) ToEntity() users.User {
	return users.User{
		Login:    dto.Login,
		Password: dto.Password,
	}
}
