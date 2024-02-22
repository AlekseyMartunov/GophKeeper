package handlers

import "GophKeeper/internal/server/entity/users"

type userDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (u *userDTO) ToEntity() users.User {
	return users.User{
		Login:    u.Login,
		Password: u.Password,
	}
}
