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
		if err := rows.Scan(&MyCard.Words.Word[count]); err != nil {
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
