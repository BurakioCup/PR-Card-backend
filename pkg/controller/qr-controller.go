package controller

import (
	"PR-Card_backend/pkg/model/dao"
	"PR-Card_backend/pkg/view"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ReadQRHandler()gin.HandlerFunc{
	return func(c *gin.Context) {
		userID := c.GetHeader("userID")
		if userID==""{
			log.Println("[ERROR] userID is empty")
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"InternalServerError",
				"userID is empty",
			)
			return
		}
		client := dao.MakeReadQRClient()
		cardPath,err := client.Request(userID)
		if err!=nil{
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Failed to select card",
			)
			return
		}
		c.JSON(http.StatusOK,view.ReturnReadQRResponse(cardPath))
	}
}
