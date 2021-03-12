package controller

import (
	"PR-Card_backend/pkg/hash"
	"PR-Card_backend/pkg/server/model/dao"
	"PR-Card_backend/pkg/server/view"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SignUpHandler()gin.HandlerFunc{
	return func(c *gin.Context) {
		userID := c.GetHeader("userID")
		if userID==""{
			log.Println("[ERROR] userID is empty")
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"UserID is empty",
			)
			return
		}
		pass := c.GetHeader("pass")
		if pass==""{
			log.Println("[ERROR] pass is empty")
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"pass is empty",
			)
			return
		}
		client := dao.MakeSignUpClient()
		loginID, err := client.Request(userID,hash.CreateHashString(pass))
		if err!=nil{
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Failed to insert user info",
			)
			return
		}
		//token,err :=jwt.CreateToken(userID)
		c.JSON(http.StatusOK, view.ReturnSignUpResponse(userID,loginID))
	}
}

func SignInHandler()gin.HandlerFunc {
	return func(c *gin.Context) {
		loginID := c.GetHeader("loginId")
		if loginID == "" {
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
		if err != nil {
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Failed to update loginID",
			)
			return
		}
		c.JSON(http.StatusOK, view.ReturnSignInResponse(loginID))
	}
}
