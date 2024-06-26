package cardhandlers

import (
	"GophKeeper/internal/entity/card"
	"time"
)

type cardName struct {
	Name string `json:"name"`
}

type cardDTO struct {
	CardName    string    `json:"card_name"`
	Number      string    `json:"number"`
	Owner       string    `json:"owner"`
	Date        string    `json:"date"`
	CreatedTime time.Time `json:"created_time"`
	CVV         string    `json:"cvv"`
	userID      int
}

func (dto *cardDTO) toEntity() card.Card {
	c := card.Card{
		Name:        dto.CardName,
		Number:      dto.Number,
		Owner:       dto.Owner,
		CVV:         dto.CVV,
		Date:        dto.Date,
		CreatedTime: dto.CreatedTime,
		UserID:      dto.userID,
	}

	return c
}

func (dto *cardDTO) fromEntity(c card.Card) {
	dto.CardName = c.Name
	dto.Owner = c.Owner
	dto.CVV = c.CVV
	dto.Number = c.Number
	dto.Date = c.Date
	dto.CreatedTime = c.CreatedTime
}

func arrDTO(c []card.Card) []cardDTO {
	res := make([]cardDTO, len(c))
	for ind, val := range c {
		res[ind].CardName = val.Name
		res[ind].Number = val.Number
		res[ind].Owner = val.Owner
		res[ind].Date = val.Date
		res[ind].CVV = val.CVV
		res[ind].CreatedTime = val.CreatedTime
	}
	return res
}
