package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/davecgh/go-spew/spew"
	"log"
)

// Token 作成関数
func CreateToken(userId string) (string, error) {
	var err error

	// 鍵となる文字列(多分なんでもいい)
	secret := "secret"

	// Token を作成
	// jwt -> JSON Web Token - JSON をセキュアにやり取りするための仕様
	// jwtの構造 -> {Base64 encoded Header}.{Base64 encoded Payload}.{Signature}
	// HS254 -> 証明生成用(https://ja.wikipedia.org/wiki/JSON_Web_Token)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"iss":   "pr-card", // JWT の発行者が入る(文字列は任意)
	})
	//Dumpを吐く
	spew.Dump(token)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}