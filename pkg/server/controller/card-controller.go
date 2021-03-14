package controller

import (
	"PR-Card_backend/pkg/server/model/dao"
	"PR-Card_backend/pkg/server/model/dto"
	"PR-Card_backend/pkg/server/view"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReadCardHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "")
	}
}

func ReadAllHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("userID")
		if userID == "" {
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
		cards, err := client.Request(userID)
		if err != nil {
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

func ReadMycardHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "")
	}
}

func CreateCardOverview() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetString("userID")
		if userID == "" {
			log.Println("[ERROR] userID is empty")
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"InternalServerError",
				"userID is empty",
			)
			return
		}

		//リクエストボディを取得
		var upr dto.RequestCardOver
		if err := c.BindJSON(&upr); err != nil {
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"RequestBody is empty",
			)
			return
		}
		clientCard := dao.MakePostCardClientClient()
		clientChart := dao.MakePostChartClientClient()

		for _, i := range upr.Status {
			err := clientChart.Request(i.ItemName, i.ItemScore)
			if err != nil {
				log.Println(err)
				view.ReturnErrorResponse(
					c,
					http.StatusInternalServerError,
					"Internal Server Error",
					"Failed to chart info",
				)
				return
			}
		}

		err := clientCard.Request(userID, upr.UserName, upr.FaceImage)
		if err != nil {
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Failed to card info",
			)
			return
		}

		c.JSON(http.StatusOK, "hey guys")
	}
}

func CreateCardDetails() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "")
	}
}

func UpdateCard() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "")
	}
}
