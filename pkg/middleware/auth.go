package middleware

import (
	"PR-Card_backend/pkg/server/view"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authenticate ユーザ認証を行ってContextへユーザID情報を保存する
func Authenticate(ginNextMethod gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			log.Println("[ERROR] token is empty")
			view.ReturnErrorResponse(
				c,
				http.StatusBadRequest,
				"Bad Request",
				"token is empty",
			)
		}
		//TODO jwtの認証を実装

		// ユーザIDをContextへ保存して以降の処理に利用する
		c.Set("userID", token)
		ginNextMethod(c)
	}
}
