package dao

import (
	"PR-Card_backend/pkg/server/model/dto"
	"database/sql"
	"errors"
	"log"
)

const(
	SelectMyCardID = "SELECT `card_id` FROM `users` WHERE `id` =?;"
	SelectMyCard = "SELECT `name_image`,`face_image`,`status_image`,`tag_image`,`free_image` FROM `cards` WHERE `id` = ?;"
	ReadAllCardsID = "SELECT `card_id` FROM `owned_cards` WHERE `user_id` = ? "
	readAllCards = "SELECT `name`,`face_image` FROM `cards` WHERE `id` = ? ;"
	UpdatCardDetailInfo= "UPDATE `cards` SET `name_image`= ?,`tag_image`=?,`free_image`=? WHERE `id`=?;"
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
	getCardInfo(cardID)
  	return MyCard, err
}



type readCardID struct{
}

func MakeReadCardIDClient()readCardID{
	return readCardID{}
}

func (infom* readCardID)Request(cardID string)(dto.MyCard,error){
	err := getCardInfo(cardID)
	if err != nil {
		return MyCard,err
	}
  	return MyCard, err
}

func getCardInfo(cardID string)error{
	row := Conn.QueryRow(SelectMyCard, cardID)
	if err := row.Scan(&MyCard.NameImage,&MyCard.FaceImage,&MyCard.StatusImage,&MyCard.TagImage,&MyCard.FreeImage); err != nil {
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

func MakeReadAllClient() raedAll {
	return raedAll{}
}

func (info *raedAll) Request(userID string) ([]dto.Card, error) {
	err := getListCardIDs(userID)
	if err != nil {
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
		Card, err = dto.ConvertToCard(rows)
		if err != nil {
			return err
		}
		Cards = append(Cards, *Card)
	}
	return err
}

func getCards() error {
	for i := 0; i < len(Cards); i++ {
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

type CreateDetail struct {
}

func MakeCreateDetailClient() CreateDetail {
	return CreateDetail{}
}

func (info *CreateDetail) Request(userID,nameImage,tagImage,freeImage string) error {
	var cardID string
	var err error
	row := Conn.QueryRow(SelectMyCardID, userID)
	if err = row.Scan(&cardID); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("Not created cards")
		}
		log.Println(err)
		return err
	}
	stmt, err := Conn.Prepare(UpdatCardDetailInfo)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(nameImage,tagImage,freeImage,cardID)
	return err
}
