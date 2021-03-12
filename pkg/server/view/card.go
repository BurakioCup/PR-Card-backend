package view

import (
	"PR-Card_backend/pkg/server/model/dto"
	"github.com/gin-gonic/gin"
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

type Error struct {
	Code		int		`json:"code"`
	Message		string	`json:"message"`
	Description	string	`json:"description"`
}

func ReturnErrorResponse(c *gin.Context, code int, msg, desc string) {
	body := Error{
		Code: code,
		Message: msg,
		Description: desc,
	}
	c.JSON(code, body)
}
