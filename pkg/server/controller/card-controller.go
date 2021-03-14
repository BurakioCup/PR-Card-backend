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
		if err!=nil{
			log.Println(err)
      view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
        "Failed to get MyCard info",
			)
			return
		}
		c.JSON(http.StatusOK, view.ReturnReadMyCardResponse(myCard))
    }
}
      
func ReadCardIDHandler()gin.HandlerFunc{
	return func(c *gin.Context) {
		cardID := c.Query("cardID")
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

func ReadAllHandler()gin.HandlerFunc{
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
		client := dao.MakeReadAllClient()
		cards,err := client.Request(userID)
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
		c.JSON(http.StatusOK, view.ReturnReadAllResponse(&cards))
	}
}

func ReadCardHandler()gin.HandlerFunc{
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
