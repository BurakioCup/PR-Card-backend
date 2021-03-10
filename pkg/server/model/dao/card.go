package dao

import (
	"PR-Card_backend/pkg/server/model/dto"
)

const(
	readAllCardsID = "SELECT `card_id` FROM `owned_cards` WHERE `user_id`=? "
	readAllCards = "SELECT * FROM `cards` WHERE ;"
)

var (
	Cards []dto.Card
	Card *dto.Card
)

// read/all
type raedAll struct {
}

func MakeReadQRClient () raedAll {
	return raedAll{}
}

func (info *raedAll)Request(userID string)([]dto.Card,error){
	err := getListCardIDs(userID)
	if err !=nil {
		return nil, err
	}

	return Cards, nil
}

func getListCardIDs(userID string)(error){
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