package dao

import (
	"github.com/google/uuid"
	"log"
)

const (
	InsertcardInfoQuery   = "INSERT INTO `users` (id,name,image_path,free_text) VALUES (?,?,?,?)"
	InsertmatrixsInfoQuery   = "INSERT INTO `users` (id,card_id,item,score) VALUES (?,?,?,?)"
)


type createChart struct {
}

func MakePostChartClientClient() createChart{
	return createChart{}
}

func (info *createChart)Request(id,item string, score int) error{
	cardID, err := uuid.NewRandom()
	if err != nil {
		log.Println("cardID generate is failed")
	}
	stmt, err := Conn.Prepare(InsertmatrixsInfoQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id,cardID,item,score)
	return  err
}



type createCard struct {
}

func MakePostCardClientClient() createCard {
	return createCard{}
}

func (info *createCard)Request(user_id,name, image_path, free_text string) error{
	stmt, err := Conn.Prepare(InsertcardInfoQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user_id,name,image_path, free_text)
	return  err
}

