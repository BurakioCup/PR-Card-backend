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
	//アカウント作成
	Server.POST("/auth/create",)
	//ユーザが持っている名刺一覧
	Server.GET("/card/list", )
	//自分の名刺を編集
	Server.GET("/card/fix", )
	//自分の名刺を登録
	Server.POST("/card/regist",)
	//名刺を表示するためのQRのリンクを送信
	Server.GET("/qr/display",)

}
