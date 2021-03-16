package controller

import (
	"PR-Card_backend/pkg/hash"
	"PR-Card_backend/pkg/server/model/dao"
	"PR-Card_backend/pkg/server/view"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GitStatusHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		gitName := c.GetHeader("gitName")
		if gitName == "" {
			log.Println("[ERROR] git account name is empty")
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"git account name is empty",
			)
			return
		}
		client := dao.MakeGitStatusClient()
		statusImage, err := client.Request(userID, hash.CreateHashString(pass))
		if err != nil {
			log.Println(err)
			view.ReturnErrorResponse(
				c,
				http.StatusInternalServerError,
				"Internal Server Error",
				"Failed to insert user info",
			)
			return
		}
	}
}