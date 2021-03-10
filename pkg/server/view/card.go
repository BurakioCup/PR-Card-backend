package view

import (
	"PR-Card_backend/pkg/server/model/dto"
	"github.com/gin-gonic/gin"
)

type ReadAllResponse struct{
	Cards *[]dto.Card `json:"cards"`
}

func ReturnReadAllResponse(cards *[]dto.Card) ReadAllResponse {
	return ReadAllResponse{Cards: cards}
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
