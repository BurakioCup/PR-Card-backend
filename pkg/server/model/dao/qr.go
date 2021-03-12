package dao

import (
	"database/sql"
	"errors"
	"log"
)

const(
	SelectUserInfoQuery = "SELECT `card_id` FROM `users` WHERE `id` = ? ;"
	SelectCardQR = "SELECT `card_image` FROM `card_qrs` WHERE `card_id` = ?; "
)

type raedQR struct {
}

func MakeReadQRClient () raedQR {
	return raedQR{}
}

func (info *raedQR)Request(userID string)(string,error){
	var cardID string
	var cardQR string
	var err error
	row := Conn.QueryRow(SelectUserInfoQuery, userID)
	if err = row.Scan(&cardID); err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("Not created cards")
		}
		log.Println(err)
		return "", err
	}
	row = Conn.QueryRow(SelectCardQR, cardID)
	if err = row.Scan(&cardQR); err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("Not created cards")
		}
		log.Println(err)
		return "", err
	}
	return cardQR,err
}
