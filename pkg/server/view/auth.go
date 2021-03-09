package view

import "github.com/gin-gonic/gin"

type SignInResponse struct{
	LoginId string `json:"loginID"`
}

func ReturnSignInResponse(loginID string) SignInResponse {
	body := SignInResponse{
		LoginId: loginID,
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