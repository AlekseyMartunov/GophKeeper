package userhandlers

import (
	"GophKeeper/internal/entity/users"
)

type userDTO struct {
	UserLogin    string `json:"login"`
	UserPassword string `json:"password"`
	NameForToken string `json:"name,omitempty"`
}

type tokenDTO struct {
	Token string `json:"token"`
}

func (dto *userDTO) ToEntity() users.User {
	return users.User{
		Login:    dto.UserLogin,
		Password: dto.UserPassword,
	}
}
