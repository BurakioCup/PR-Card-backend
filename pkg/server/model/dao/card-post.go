package dao

import "github.com/google/uuid"

const (
	InsertcardInfoQuery    = "INSERT INTO `cards` (id,name,face_image) VALUES (?,?,?)"
	InsertmatrixsInfoQuery = "INSERT INTO `card_matrixs` (id,card_id,item,socre) VALUES (?,?,?,?)"
)

type createChart struct {
}

func MakePostChartClientClient() createChart {
	return createChart{}
}

func (info *createChart) Request(item string, socre int) error {

	/*cardID, err := uuid.NewRandom()
	if err != nil {
		log.Println("cardID generate is failed")
	}
	*/
	//オートインサートにしませんか？
	id := uuid.New()
	cardID := "124"

	stmt, err := Conn.Prepare(InsertmatrixsInfoQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id, cardID, item, socre)
	return err
}

type createCard struct {
}

func MakePostCardClientClient() createCard {
	return createCard{}
}

func (info *createCard) Request(user_id, name, image_path string) error {

	stmt, err := Conn.Prepare(InsertcardInfoQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user_id, name, image_path)
	return err
}
