package controller

import (
	"PR-Card_backend/pkg/jwt"
	"PR-Card_backend/pkg/server/model/dao"
	"PR-Card_backend/pkg/server/view"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func SignUpHandler()gin.HandlerFunc{
	return func(c *gin.Context) {
		userID := c.GetHeader("userId")
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
			log.Println("[ERROR] userID is empty")
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"UserID is empty",
			)
			return
		}
		client := dao.MakeSignUpClient()
		//loginId, err := client.Request(userID,hash.CreateHashString(pass))
		loginId, err := client.Request(userID,pass)
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
		token,err :=jwt.CreateToken(userID)
		c.JSON(http.StatusOK, view.ReturnSignUpResponse(token,loginId))
	}
}

func SigninHandler()gin.HandlerFunc{
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "")
	}
}
