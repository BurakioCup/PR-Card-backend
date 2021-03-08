package controller

import (
	"PR-Card_backend/pkg/hash"
	"PR-Card_backend/pkg/model/dao"
	"PR-Card_backend/pkg/view"
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
		tokenId, err := client.Request(userID,hash.CreateHashString(pass))
		if err!=nil{
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Failed to get the list of articles",
			)
			return
		}
		c.JSON(http.StatusOK, view.ReturnSignUpResponse(tokenId,loginId))
	}
}

func SigninHandler()gin.HandlerFunc{
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, "")
	}
}
