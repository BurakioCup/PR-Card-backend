package server

import (
	"PR-Card_backend/pkg/middleware"
	"PR-Card_backend/pkg/server/controller"
	"github.com/gin-gonic/gin"
)

var (
	//Server gin flameworkのserver
	Server *gin.Engine
)

func init() {
	Server = gin.Default()
	//アカウント作成
	Server.POST("/sign/", controller.CreateAuthHandler())
	//アカウント認証
	Server.GET("/sign/in", controller.SigninHandler())
	//ユーザが持っている名刺一覧
	Server.GET("/read/all", middleware.Authenticate(controller.ReadAllHandler()))
	//一覧から一つの名刺を詳細表示
	Server.GET("/read/card", middleware.Authenticate(controller.ReadCardHandler()))
	//自分の名刺を編集するために最初の状態を送信
	Server.GET("/read/mycard", middleware.Authenticate(controller.ReadMycardHandler()))
	//自分の変更後の名刺を登録
	Server.POST("/create/card", middleware.Authenticate(controller.CreateCard()))
	//自分の変更後の名刺を登録
	Server.PUT("/update/card", middleware.Authenticate(controller.UpdateCard()))
	//名刺を表示するためのQRのリンクを送信
	Server.GET("/read/qr", middleware.Authenticate(controller.ReadQR()))
}
