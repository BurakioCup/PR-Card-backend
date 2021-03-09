package controller

import (
	"PR-Card_backend/pkg/model/dao"
	"PR-Card_backend/pkg/view"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateAuthHandler()gin.HandlerFunc{
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "")
	}
}

func SignInHandler()gin.HandlerFunc{
	return func(c *gin.Context) {
		loginID := c.GetHeader("loginId")
		if loginID==""{
			log.Println("[ERROR] loginID is empty")
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"loginID is empty",
			)
			return
		}
		client := dao.MakeSignInClient()
		loginID, err := client.Request(loginID)
		if err!=nil{
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Failed to loginID update",
			)
			return
		}
		c.JSON(http.StatusOK, view.ReturnSignInResponse(loginID))
	}
}
