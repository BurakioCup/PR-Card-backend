package dao

import (
	"PR-Card_backend/pkg/server/model/dto"
	"database/sql"
	"errors"
	"log"
)

const(
	SelectMyCardID = "SELECT `card_id` FROM `users` WHERE `id` =?;"
	SelectMyCard = "SELECT `name`,`name_image`,`face_image`,`status_image`,`tag_image`,`free_image` FROM `cards` WHERE `id` = ?;"
	SelectOtherCard = "SELECT `name_image`,`face_image`,`status_image`,`tag_image`,`free_image` FROM `cards` WHERE `id` = ?;"
	ReadAllCardsID = "SELECT `card_id` FROM `owned_cards` WHERE `user_id` = ? "
	readAllCards = "SELECT `name`,`face_image` FROM `cards` WHERE `id` = ? ;"
	UpdateCardDetailInfo = "UPDATE `cards` SET `name`=?,`name_image`= ?,`tag_image`=?,`free_image`=? WHERE `id`=?;"
	)

type readMyCard struct{
}

func MakeReadMyCardClient()readMyCard{
	return readMyCard{}
}

func (info *readMyCard)Request(userID string)(dto.MyCard,error){
	var myCard dto.MyCard
	var cardID string
	var err error
	row := Conn.QueryRow(SelectMyCardID, userID)
	if err = row.Scan(&cardID); err != nil {
		if err == sql.ErrNoRows {
			return myCard, errors.New("Not created cards")
		}
		log.Println(err)
		return myCard, err
	}
	myCard,err=getMyCardInfo(cardID)
	if err != nil {
		log.Println(err)
		return myCard,err
	}
  	return myCard, err
}

func getMyCardInfo(cardID string)(dto.MyCard,error){
	var myCard dto.MyCard
	row := Conn.QueryRow(SelectMyCard, cardID)
	if err := row.Scan(&myCard.UserName,&myCard.NameImage,&myCard.FaceImage,&myCard.StatusImage,&myCard.TagImage,&myCard.FreeImage); err != nil {
		if err == sql.ErrNoRows {
			return myCard,errors.New("Not created cards")
		}
		log.Println(err)
		return myCard,err
	}
	return myCard,nil
}

type readCardID struct{
}

func MakeReadCardIDClient()readCardID{
	return readCardID{}
}

func (infom* readCardID)Request(cardID string)(dto.DetailCard,error){
	detailCard,err := getOtherCardInfo(cardID)
	if err != nil {
		return detailCard,err
	}
  	return detailCard, err
}

func getOtherCardInfo(cardID string)(dto.DetailCard,error){
	var card dto.DetailCard
	row := Conn.QueryRow(SelectOtherCard, cardID)
	if err := row.Scan(&card.NameImage,&card.FaceImage,&card.StatusImage,&card.TagImage,&card.FreeImage); err != nil {
		if err == sql.ErrNoRows {
			return card,errors.New("Not created cards")
		}
		log.Println(err)
		return card,err
	}
	return card,nil
}

// read/all
type raedAll struct {
}

func MakeReadAllClient() raedAll {
	return raedAll{}
}

func (info *raedAll) Request(userID string) ([]dto.Card, error) {
	var Cards []dto.Card
	Cards,err := getListCardIDs(userID)
	if err != nil {
		return nil, err
	}
	err = getCards(Cards)
	return Cards, nil
}

func getListCardIDs(userID string)([]dto.Card,error){
	var cards []dto.Card
	rows, err := Conn.Query(ReadAllCardsID, userID)
	if err != nil {
		return nil,err
	}
	defer rows.Close()
	//取得してきた複数(単数)のレコード1つずつ処理
	for rows.Next() {
		Card, err := dto.ConvertToCard(rows)
		if err != nil {
			return nil,err
		}
		cards = append(cards, *Card)
	}
	return cards,err
}

func getCards(Cards []dto.Card) error {
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

func (info *CreateDetail) Request(userID,name,nameImage,tagImage,freeImage string) error {
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
	stmt, err := Conn.Prepare(UpdateCardDetailInfo)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(name,nameImage,tagImage,freeImage,cardID)
	return err
}
