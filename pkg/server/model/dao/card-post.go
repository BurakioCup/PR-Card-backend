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

func (info *createChart) Request(face_image, status_image string) error {

	id := uuid.New()

	stmt, err := Conn.Prepare(InsertcardInfoQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id, face_image, status_image)
	return err
}
