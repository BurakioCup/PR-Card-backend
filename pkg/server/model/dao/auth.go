package dao

import (
	"fmt"
	"github.com/google/uuid"
	"log"
)

const(
	InsertUserInfoQuery="INSERT INTO `users` (id,pass,login_id) VALUES (?,?,?)"
)
// sign/up
type signUp struct {
}

func MakeSignUpClient () signUp {
	return signUp{}
}

func (info *signUp)Request(userID,pass string)(string,error){
	tokenId, err := uuid.NewRandom()
	if err != nil {
		log.Println("tokenID generate is failed")
	}
	stmt, err := Conn.Prepare(InsertUserInfoQuery)
	if err != nil {
		return "",err
	}
	fmt.Println(pass)
	_, err = stmt.Exec(userID, pass, tokenId)
	return tokenId.String(),err
}
