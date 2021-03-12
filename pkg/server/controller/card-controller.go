package controller

import (
	"PR-Card_backend/pkg/server/model/dao"
	"PR-Card_backend/pkg/server/view"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ReadCardIDHandler()gin.HandlerFunc{
	return func(c *gin.Context) {
		cardID := c.Param("cardID")
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
		fmt.Println(card)
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
