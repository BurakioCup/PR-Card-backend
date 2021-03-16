package util

import (
	"PR-Card_backend/pkg/server/model/dao"
	"database/sql"
	"errors"
	"log"
)

const (
	SelectMyCardID = "SELECT `card_id` FROM `users` WHERE `id` =?;"
)

func GetCardID(userID string) (string,error){
	var cardID string
	var err error
	row := dao.Conn.QueryRow(SelectMyCardID, userID)
	if err = row.Scan(&cardID); err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("Not created cards")
		}
		log.Println(err)
		return  "",err
	}
	return cardID,err
}
