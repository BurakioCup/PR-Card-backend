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
	Server.POST("/sign/up", controller.SignUpHandler())
	//アカウント認証
	Server.GET("/sign/in", controller.SignInHandler())
	//ユーザが持っている名刺一覧
	Server.GET("/read/all", middleware.Authenticate(controller.ReadAllHandler()))
	//一覧から一つの名刺を詳細表示
	Server.GET("/read", middleware.Authenticate(controller.ReadCardIDHandler()))
	//自分の名刺を編集するために最初の状態を送信
	Server.GET("/read/myCard", middleware.Authenticate(controller.ReadMyCardHandler()))
    //名前・顔写真URLをパラメーターとってたかしにぶん投げて返却されたものを保存
	Server.POST("/create/card/overview", middleware.Authenticate(controller.CreateCardOverview()))
	//グラフのデータ、自分を表す言葉、自由記述欄を受け取りDBに保管
	Server.POST("/create/card/details", middleware.Authenticate(controller.CreateCardDetails()))
	//自分の変更後の名刺を登録
	Server.PUT("/update/card", middleware.Authenticate(controller.UpdateCard()))
	//名刺を表示するためのQRのリンクを送信
	Server.GET("/read/qr", middleware.Authenticate(controller.ReadQRHandler()))
}
