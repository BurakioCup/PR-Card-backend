package dao

import (
	"PR-Card_backend/pkg/server/model/dto"
	"database/sql"
	"errors"
	"log"
)

const(
	SelectMyCardID = "SELECT `card_id` FROM `users` WHERE `id` =?;"
	SelectMyCard = "SELECT `name`,`nick_name`,`face_image`,`status_image`,`free_text` FROM `cards` WHERE `id` = ?;"
	SelectMyCardWords = "SELECT `word` FROM `card_my_word` WHERE `card_id` = ?"
	ReadAllCardsID = "SELECT `card_id` FROM `owned_cards` WHERE `user_id` = ? "
	readAllCards = "SELECT `name`,`face_image` FROM `cards` WHERE `id` = ? ;"
	)



var (
	Cards []dto.Card
	Card *dto.Card
	MyCard dto.MyCard
)

type readMyCard struct{
}

func MakeReadMyCardClient()readMyCard{
	return readMyCard{}
}

func (info *readMyCard)Request(userID string)(dto.MyCard,error){
	var cardID string
	var err error
	row := Conn.QueryRow(SelectMyCardID, userID)
	if err = row.Scan(&cardID); err != nil {
		if err == sql.ErrNoRows {
			return MyCard, errors.New("Not created cards")
		}
		log.Println(err)
		return MyCard, err
	}
	getMyCardInfo(cardID)
	getMyCardWord(cardID)
	return MyCard, err
}

func getMyCardWord(cardID string)error{
	count := 0
	rows,err := Conn.Query(SelectMyCardWords, cardID)
	if err != nil {
		return err
	}
	defer rows.Close()
	//取得してきた複数(単数)のレコード1つずつ処理
	for rows.Next() {
		//レコードを構造体Articleに整形
		if err := rows.Scan(&MyCard.Words[count]); err != nil {
			if err == sql.ErrNoRows {
				return  nil
			}
			log.Println(err)
			return  err
		}
		count++
	}
	return nil
}

func getMyCardInfo(cardID string)error{
	row := Conn.QueryRow(SelectMyCard, cardID)
	if err := row.Scan(&MyCard.UserName,&MyCard.NickName,&MyCard.FaceImage,&MyCard.StatusImage,&MyCard.FreeText); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("Not created cards")
		}
		log.Println(err)
		return err
	}
	return nil
}

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
	rows, err := Conn.Query(ReadAllCardsID, userID)
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