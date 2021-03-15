package dao

import (
	"database/sql"
	"errors"
	"log"
)

const (
	InsertcardInfoQuery    = "INSERT INTO `cards` (id,face_image,status_image) VALUES (?,?,?)"
	InsertCardDetailInfo   = "UPDATE `cards` SET `name_image`= ?,`tag_image`=?,`free_image`=? WHERE `id`=?;"
)

type createChart struct {
}

func MakePostChartClientClient() createChart {
	return createChart{}
}

func (info *createChart) Request(id,face_image, status_image string) error {


	stmt, err := Conn.Prepare(InsertcardInfoQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id, face_image, status_image)
	return err
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
	stmt, err := Conn.Prepare(InsertCardDetailInfo)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(nameImage,tagImage,freeImage,cardID)
	return err
}