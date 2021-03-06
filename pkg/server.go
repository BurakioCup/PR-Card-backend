package pkg

import (
	//"PR-Card_backend/pkg/controller"
	"github.com/gin-gonic/gin"
)

var (
	//Server gin flameworkのserver
	Server *gin.Engine
)

func init() {
	Server = gin.Default()
	//
	Server.GET("/read/articles", )

}
