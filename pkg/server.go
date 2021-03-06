package pkg

import (
	"PR-Card_backend/pkg/controller"
	"github.com/gin-gonic/gin"
)

var (
	//Server gin flameworkのserver
	Server *gin.Engine
)

func init() {
	Server = gin.Default()
	//アカウント作成
	Server.POST("/create/auth", controller.CreateAuthHandler())
	//ユーザが持っている名刺一覧
	Server.GET("/read/cards", controller.ReadCardsHandler())
	//一覧から一つの名刺を詳細表示
	Server.GET("/read/card", controller.ReadCardHandler())
	//自分の名刺を編集するために最初の状態を送信
	Server.GET("/read/mycard", controller.ReadMycardHandler())
	//自分の変更後の名刺を登録
	Server.POST("/create/card", controller.CreateCard())
	//自分の変更後の名刺を登録
	Server.PUT("/update/card", controller.UpdateCard())
	//名刺を表示するためのQRのリンクを送信
	Server.GET("/read/qr", controller.ReadQR())
}
