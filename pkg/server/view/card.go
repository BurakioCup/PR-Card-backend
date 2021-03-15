package view

import (
	"PR-Card_backend/pkg/server/model/dto"
)

type ReadAllResponse struct {
  Cards *[]dto.Card `json:"cards"`
}

type ReadCardResponse struct{
	NameImage string `json:"nameImage"`
	FaceImage string `json:"faceImage"`
	TagImage string `json:"tagImage"`
	StatusImage string `json"statusImage"`
	FreeImage string `json:"freeImage"`
}

func ReturnReadCard(card dto.MyCard)ReadCardResponse{
	return ReadCardResponse{NameImage: card.NameImage,FaceImage: card.FaceImage,TagImage: card.TagImage,
		StatusImage: card.StatusImage,FreeImage: card.FreeImage}
}




func ReturnReadAllResponse(cards *[]dto.Card) ReadAllResponse {
	return ReadAllResponse{Cards: cards}
}
