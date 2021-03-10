package dto

import (
	"database/sql"
	"log"
)

type Card struct {
	CardID string
	UserName string
	FaceImage string
}

func ConvertToCard(row *sql.Rows)(*Card, error){
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
