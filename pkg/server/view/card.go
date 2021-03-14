package view

import (
	"PR-Card_backend/pkg/server/model/dto"
)

type ReadMyCardResponse struct{
	UserName string `json:"userName"`
	FaceImage string `json:"faceImage"`
	NickName string `json:"nickName"`
	StatusImage string `json"statusImage"`
	Words [4]string `json:"words"`
	FreeText string `json:"freeText"`
}

type ReadCardResponse struct{
	UserName string `json:"userName"`
	FaceImage string `json:"faceImage"`
	NickName string `json:"nickName"`
	StatusImage string `json"statusImage"`
	Words [4]string `json:"words"`
	FreeText string `json:"freeText"`
}

func ReturnReadMyCardResponse(myCard dto.MyCard)ReadMyCardResponse{
	return ReadMyCardResponse{UserName: myCard.UserName,FaceImage: myCard.FaceImage,NickName: myCard.NickName,
		StatusImage: myCard.StatusImage,Words: myCard.Words,FreeText: myCard.FreeText}
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
