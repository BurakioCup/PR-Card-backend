package view

import "github.com/gin-gonic/gin"

type SignUpResponse struct{
	TokenId string `json:"tokenID"`
	LoginId string `json:"loginID"`
}

func ReturnSignUpResponse(c *gin.Context, tokenId string, uuid string)SignUpResponse{
	body :=SignUpResponse{
		TokenId: tokenId,
		LoginId: uuid,
	}
	return body
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
