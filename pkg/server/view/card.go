package view

import (
	"PR-Card_backend/pkg/server/model/dto"
)

type ReadAllResponse struct {
	Cards *[]dto.Card `json:"cards"`
}

func ReturnReadAllResponse(cards *[]dto.Card) ReadAllResponse {
	return ReadAllResponse{Cards: cards}
}
