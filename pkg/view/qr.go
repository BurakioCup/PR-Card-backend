package view

import "github.com/gin-gonic/gin"

type ReadQRResponse struct{
//	TODO json
cardImage string `json:""`
}

func ReturnReadQRResponse(imagePath string) ReadQRResponse {
	body := ReadQRResponse{
		cardImage: imagePath,
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