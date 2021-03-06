package main

import (
	"PR-Card_backend/configs"
	"PR-Card_backend/pkg"
	"PR-Card_backend/pkg/model/dao"
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
	pkg.Server.Run(addr)
}