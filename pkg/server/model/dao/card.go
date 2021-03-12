package dao

import (
	"PR-Card_backend/pkg/server/model/dto"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

const(
	readAllCardsID = "SELECT `card_id` FROM `owned_cards` WHERE `user_id` = ? "
	readAllCards = "SELECT `name`,`image_path` FROM `cards` WHERE `id` = ? ;"
)

var (
	Cards []dto.Card
	Card *dto.Card
)

// read/all
type raedAll struct {
}

func MakeReadAllClient () raedAll {
	return raedAll{}
}

func (info *raedAll)Request(userID string)([]dto.Card,error){
	err := getListCardIDs(userID)
	if err !=nil {
		return nil, err
	}
	err = getCards()
	return Cards, nil
}

func getListCardIDs(userID string)error{
	rows, err := Conn.Query(readAllCardsID, userID)
	if err != nil {
		return err
	}
	defer rows.Close()
	//取得してきた複数(単数)のレコード1つずつ処理
	for rows.Next() {
		//レコードを構造体Articleに整形
		Card, err = dto.ConvertToCard(rows)
		if err != nil {
			return err
		}
		Cards = append(Cards, *Card)
	}
	return err
}

func getCards()error{
	for i:=0; i<len(Cards); i++{
		row := Conn.QueryRow(readAllCards, Cards[i].CardID)
		fmt.Println("aa")
		if err := row.Scan(&Cards[i].UserName,&Cards[i].FaceImage); err != nil {
			if err == sql.ErrNoRows {
				return errors.New("Faild get cards info")
			}
			log.Println(err)
			return err
		}
	}
	return nil
}