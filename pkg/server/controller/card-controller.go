package controller

import (
	"PR-Card_backend/pkg/server/model/dao"
	"PR-Card_backend/pkg/server/model/dto"
	"PR-Card_backend/pkg/server/view"
	"bytes"

	//"bytes"
	"fmt"
	"github.com/google/uuid"

	//"bytes"
	"encoding/json"
	//"io/ioutil"
	"log"
	"net/http"
	//"net/http/cookiejar"

	"github.com/gin-gonic/gin"
)

func ReadMyCardHandler() gin.HandlerFunc {
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
		client := dao.MakeReadMyCardClient()
		myCard, err := client.Request(userID)
		if err != nil {
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

func ReadCardIDHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		cardID := c.Query("cardID")
		if cardID == "" {
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
		card, err := client.Request(cardID)
		if err != nil {
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

func ReadCardHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "")
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
		cardID := uuid.New().String()
		req:=dto.RequestCardResponse(cardID,upr.FaceImage,upr.Status)
		//takashi serverへのPOST処理
		endpoint := "http://localhost:3000/newIconChart"

		b, _ := json.Marshal(req)
		fmt.Println("aa")
		reqBody, err := http.NewRequest(
			"POST",
			endpoint,
			bytes.NewBuffer(b),
		)
		fmt.Println("bb")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("aabb")
		reqBody.Header.Set("Content-Type", "application/json")
		fmt.Println("aacc")
		//client := new(http.Client)
		client := &http.Client{}
		fmt.Println(reqBody)
		resp, err := client.Do(reqBody)
		fmt.Println("aa")
		if err != nil {
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"RequestBody is empty",
			)
			return
		}
		fmt.Println("あ")
		//ボディの取得
		var requestBody view.CreateCardOverResponse
		if err := json.NewDecoder(resp.Body).Decode(&requestBody); err != nil {
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"RequestBody is empty",
			)
			return
		}
		fmt.Println(requestBody)


		responseBody := view.ReturnCreateCardResponse(requestBody.FaceImage,requestBody.StatusImage)
		clientCard := dao.MakePostChartClientClient()

		_ = clientCard.Request(cardID, responseBody.FaceImage, responseBody.StatusImage)
		if err != nil {
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Failed to get MyCard info",
			)
			return
		}
		c.JSON(http.StatusOK, view.ReturnCreateCardResponse(responseBody.FaceImage,responseBody.StatusImage))
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
