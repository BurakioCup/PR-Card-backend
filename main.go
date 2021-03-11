package main

import (
	"PR-Card_backend/configs"
	"PR-Card_backend/pkg/server"
	"PR-Card_backend/pkg/server/model/dao"
	"log"
)

func main() {
	//DBのコネクションを確率
	err := dao.Init()
	if err != nil {
		log.Fatal(err)
	}

	//サーバを起動
	addr := configs.GetServerPort()
	server.Server.Run(addr)
}