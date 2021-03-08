package hash

import (
	"golang.org/x/crypto/bcrypt"
)

func CreateHashString(str string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(str), 10)
	//パスワードを検証する場合は下記のコマンドを使用
	//https://zenn.dev/kou_pg_0131/articles/go-digest-and-compare-by-bcrypt
	//err := bcrypt.CompareHashAndPassword(hashed, password)
	return string(hashed)
}