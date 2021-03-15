package dao

import "github.com/google/uuid"

const (
	InsertcardInfoQuery    = "INSERT INTO `cards` (id,face_image,status_image) VALUES (?,?,?)"
)

type createChart struct {
}

func MakePostChartClientClient() createChart {
	return createChart{}
}

func (info *createChart) Request(item string, socre int) error {

	id := uuid.New()

	stmt, err := Conn.Prepare(InsertcardInfoQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id, item, socre)
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
