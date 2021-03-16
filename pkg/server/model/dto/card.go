package dto

import (
	"database/sql"
	"log"
)

type DetailCard struct{
	NameImage string
	FaceImage string
	StatusImage string
	TagImage string
	FreeImage string
}


type MyCard struct{
	UserName string
	NameImage string
	FaceImage string
	StatusImage string
	TagImage string
	FreeImage string
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
