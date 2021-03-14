package dto

import (
	"database/sql"
	"log"
)

type MyCard struct{
	UserName string
	FaceImage string
	NickName string
	StatusImage string
	Words [4]string
	FreeText string
}



type Card struct {
	CardID    string `json:"cardID"`
	UserName  string `json:"userName"`
	FaceImage string `json:"faceImage"`
}

func ConvertToCard(row *sql.Rows) (*Card, error) {
	var card Card
	if err := row.Scan(&card.CardID); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return &card, nil
}
