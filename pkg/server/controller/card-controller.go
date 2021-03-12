package controller

import (
	"PR-Card_backend/pkg/server/model/dao"
	"PR-Card_backend/pkg/server/view"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ReadCardIDHandler()gin.HandlerFunc{
	return func(c *gin.Context) {
		//TODO 下記２行を決定
		cardID := c.Param("cardID")
		cardID=c.Query("cardID")
		//fmt.Println(a)
		if cardID==""{
			log.Println("[ERROR] cardID is empty")
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"InternalServerError",
				"cardID is empty",
			)
			return
		}
		client := dao.MakeReadCardIDClient()
		card,err := client.Request(cardID)
		if err!=nil {
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Failed to get Card infomation",
			)
			return
		}
		c.JSON(http.StatusOK, view.ReturnReadCard(card))
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
