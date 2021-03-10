package dao

import (
	"database/sql"
	"errors"
	"log"
)

const(
	InsertUserInfoQuery="SELECT `card_id` FROM `users` WHERE `id` = ? ;"
)
// sign/up
type raedQR struct {
}

func MakeReadQRClient () raedQR {
	return raedQR{}
}

func (info *raedQR)Request(userID string)(string,error){
	var cardPath string
	var err error
	row := Conn.QueryRow(InsertUserInfoQuery, userID)
	if err = row.Scan(&cardPath); err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("Not created cards")
		}
		log.Println(err)
		return "", err
	}
	return cardPath,err
}
