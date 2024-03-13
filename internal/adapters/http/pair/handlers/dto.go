package pairhandlers

import (
	"GophKeeper/internal/entity/pairs"
	"time"
)

type nameDTO struct {
	Name string `json:"name"`
}

type pairDTO struct {
	Login       string    `json:"login"`
	Password    string    `json:"password"`
	Name        string    `json:"name"`
	UserID      int       `json:"-"`
	CreatedTime time.Time `json:"created_time,omitempty"`
}

func (dto *pairDTO) toEntity() pairs.Pair {
	return pairs.Pair{
		Login:    dto.Login,
		Password: dto.Password,
		Name:     dto.Name,
		UserID:   dto.UserID,
	}
}

func (dto *pairDTO) fromEntity(pair pairs.Pair) {
	dto.Password = pair.Password
	dto.Login = pair.Login
	dto.Name = pair.Name
	dto.CreatedTime = pair.CreatedTime
}

func arrDTO(p []pairs.Pair) []pairDTO {
	res := make([]pairDTO, len(p))
	for ind, val := range p {
		res[ind].Name = val.Name
		res[ind].Login = val.Login
		res[ind].Password = val.Password
		res[ind].CreatedTime = val.CreatedTime
	}
	return res
}
