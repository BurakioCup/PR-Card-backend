package view

import (
	"PR-Card_backend/pkg/server/model/dto"
)

type ReadCardResponse struct{
	UserName string `json:"userName"`
	FaceImage string `json:"faceImage"`
	NickName string `json:"nickName"`
	StatusImage string `json"statusImage"`
	Words [4]string `json:"words"`
	FreeText string `json:"freeText"`
}

func ReturnReadCard(card dto.MyCard)ReadCardResponse{
	return ReadCardResponse{UserName: card.UserName,FaceImage: card.FaceImage,NickName: card.NickName,
		StatusImage: card.StatusImage,Words: card.Words,FreeText: card.FreeText}
}

type ReadAllResponse struct{
	Cards *[]dto.Card `json:"cards"`
}

func ReturnReadAllResponse(cards *[]dto.Card) ReadAllResponse {
	return ReadAllResponse{Cards: cards}
}
