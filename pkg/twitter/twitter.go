package main

import (
	"fmt"
	"net/url"
	"gopkg.in/ini.v1"
	"github.com/ChimeraCoder/anaconda"
)

// apiの設定保存用
type ConfigList struct{
	APIKey string
	APISecretKey string
	AccessToken string
	AccessTokenSecret string
}

var Config ConfigList

// 初期処理 config.iniを読み込む
func init() {
	cfg, _ := ini.Load("config.ini")
	Config = ConfigList{
		APIKey: cfg.Section("twitter").Key("api_key").String(),
		APISecretKey: cfg.Section("twitter").Key("api_secret_key").String(),
		AccessToken: cfg.Section("twitter").Key("access_token").String(),
		AccessTokenSecret: cfg.Section("twitter").Key("access_token_secret").String(),
	}
}


func main() {
	// anacondaにconfig.iniの内容を設定
	anaconda.SetConsumerKey(Config.APIKey)
	anaconda.SetConsumerSecret(Config.APISecretKey)
	api := anaconda.NewTwitterApi(Config.AccessToken, Config.AccessTokenSecret)

	// twitter apiのパラメータを設定
	// 30件取得する
	v := url.Values{}
	v.Set("count", "10")

	// tweetを取得
	searchResult, _ := api.GetSearch("ハックツハッカソン", v)
	for _, tweet := range searchResult.Statuses {
		fmt.Println(tweet.Text)
	}
}