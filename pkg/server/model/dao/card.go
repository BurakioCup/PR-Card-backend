package dao

import (
	"PR-Card_backend/pkg/server/model/dto"
	"database/sql"
	"errors"
	"log"
)

const(
	SelectMyCardID = "SELECT `card_id` FROM `users` WHERE `id`=? "
	SelectMyCard = "SELECT "
)

type readMyCard struct{
}

func MakeReadMyCardClient()readMyCard{
	return readMyCard{}
}

var MyCard dto.MyCard

func (info *readMyCard)Request(userID string)(dto.MyCard,error){
	var cardID string
	var err error
	row := Conn.QueryRow(SelectMyCard, userID)
	if err = row.Scan(&cardID); err != nil {
		if err == sql.ErrNoRows {
			return MyCard, errors.New("Not created cards")
		}
		log.Println(err)
		return MyCard, err
	}
	getMyCardInfo(cardID)
	return MyCard, err
}

func getMyCardInfo(cardID string){
	row := Conn.QueryRow(SelectMyCard, cardID)
}
