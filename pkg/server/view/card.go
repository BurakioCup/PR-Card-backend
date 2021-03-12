package view

import (
	"PR-Card_backend/pkg/server/model/dto"
	"github.com/gin-gonic/gin"
)

type ReadMyCardResponse struct{
	UserName string `json:"userName"`
	FaceImage string `json:"faceImage"`
	NickName string `json:"nickName"`
	StatusImage string `json"statusImage"`
	Words [4]string `json:"words"`
	FreeText string `json:"freeText"`
}

//type ReadMyCardResponse struct{
//	MyCard dto.MyCard `json:"myCard"`
//}

func ReturnReadMyCardResponse(myCard dto.MyCard)ReadMyCardResponse{
	return ReadMyCardResponse{UserName: myCard.UserName,FaceImage: myCard.FaceImage,NickName: myCard.NickName,
		StatusImage: myCard.StatusImage,Words: myCard.Words,FreeText: myCard.FreeText}
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