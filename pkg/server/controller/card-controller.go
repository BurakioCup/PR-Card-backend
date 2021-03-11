package controller

import (
	"PR-Card_backend/pkg/server/model/dao"
	"PR-Card_backend/pkg/server/view"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ReadMyCardHandler()gin.HandlerFunc{
	return func(c *gin.Context) {
		userID := c.GetString("userID")
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
		client := dao.MakeReadMyCardClient()
		myCard,err := client.Request(userID)

		c.JSON(http.StatusOK, "")
	}
}

func ReadCardsHandler()gin.HandlerFunc{
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "")
	}
}

func ReadMycardHandler()gin.HandlerFunc{
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "")
	}
}

func CreateCard()gin.HandlerFunc{
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "")
	}
}

func UpdateCard()gin.HandlerFunc{
	return func(c *gin.Context) {

	c.JSON(http.StatusOK, "")
	}
}
