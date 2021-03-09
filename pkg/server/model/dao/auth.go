package dao

import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"log"
)

const(
	AuthenticationLoginID = "SELECT id FROM `users` WHERE `login_id`=? ;"
	UpdateLoginInfoQuery = "UPDATE `users` set`login_id` = ? where `id` = ? ;"
)
// sign/in
type signIn struct {
}

func MakeSignInClient () signIn {
	return signIn{}
}

func (info *signIn)Request(loginID string)(string,error){
	var userID string
	var err error
	row := Conn.QueryRow(AuthenticationLoginID, loginID)
	if err = row.Scan(&userID); err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("loginID is not true")
		}
		log.Println(err)
		return "", err
	}

	loginID = uuid.NewString()
	if err != nil {
		log.Println("tokenID generate is failed")
	}

	stmt, err := Conn.Prepare(UpdateLoginInfoQuery)
	if err != nil {
		return "",err
	}

	_, err = stmt.Exec(loginID,userID)
	return loginID,err
}