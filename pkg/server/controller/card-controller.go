package controller

import (
	"PR-Card_backend/pkg/server/model/dao"
	"PR-Card_backend/pkg/server/model/dto"
	"PR-Card_backend/pkg/server/view"
	"PR-Card_backend/pkg/util"
	"bytes"
	"github.com/google/uuid"
	"encoding/json"
	"log"
	"net/http"
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
		c.JSON(http.StatusOK, view.ReturnReadMyCard(myCard))
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
		endpoint := "https://us-central1-prcard-ae898.cloudfunctions.net/PR_card/newIconChart"

		b, _ := json.Marshal(req)

		reqBody, err := http.NewRequest(
			"POST",
			endpoint,
			bytes.NewBuffer(b),
		)

		if err != nil {
			log.Println(err)
			return
		}

		reqBody.Header.Set("Content-Type", "application/json")

		//client := new(http.Client)
		client := &http.Client{}
		resp, err := client.Do(reqBody)

		if err != nil {
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"RequestBody is empty",
			)
			return
		}
		//ボディの取得
		var requestBody view.CreateCardOverResponse
		if err := json.NewDecoder(resp.Body).Decode(&requestBody); err != nil {
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"client POST RequestBody is empty",
			)
			return
		}

		responseBody := view.ReturnCreateCardResponse(requestBody.FaceImage,requestBody.StatusImage)
		clientCard := dao.MakePostChartClientClient()

		_ = clientCard.Request(cardID, responseBody.FaceImage, responseBody.StatusImage,userID)
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
		_ = clientCard.UpdateRequest(cardID,userID)
		if err != nil {
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Not update userTable",
			)
			return
		}
		c.JSON(http.StatusOK, view.ReturnCreateCardResponse(responseBody.FaceImage,responseBody.StatusImage))
	}
}

func CreateCardDetails() gin.HandlerFunc {
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
		var req dto.CardDetailRequest
		if err := c.BindJSON(&req); err != nil {
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"Request Body is error",
			)
			return
		}
		cardID,err := util.GetCardID(userID)
		if err != nil {
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"InternalServerErro",
				"get cardID is failed",
			)
			return
		}

		reqNode:=dto.RequestCardDetailResponse(cardID,req)
		endpoint := "https://us-central1-prcard-ae898.cloudfunctions.net/PR_card/createNameTexts"
		b, _ := json.Marshal(reqNode)
		reqBody, err := http.NewRequest(
			"POST",
			endpoint,
			bytes.NewBuffer(b),
		)
		if err != nil {
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"InternalServerErro",
				"Image generate is failed",
			)
			return
		}

		reqBody.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(reqBody)
		if err != nil {
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"Error from Image generater",
			)
			return
		}

		//画像生成APIからのレスポンスの取得
		var requestBody view.CardDetailsResponse
		if err := json.NewDecoder(resp.Body).Decode(&requestBody); err != nil {
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"InternalServerError",
				"ResponseBody from Image generater is empty",
			)
			return
		}

		DBclient := dao.MakeCreateDetailClient()
		_ = DBclient.Request(userID,req.UserName.Name,requestBody.NameImage,requestBody.TagImage,requestBody.FreeImage)
		if err != nil {
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Failed to insert card info",
			)
			return
		}
		c.JSON(http.StatusOK, requestBody)

	}
}
