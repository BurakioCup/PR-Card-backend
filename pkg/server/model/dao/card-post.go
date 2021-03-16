package dao

const (
	InsertcardInfoQuery    = "INSERT INTO `cards` (id,face_image,status_image) VALUES (?,?,?)"
	UpdatUserInfo = "UPDATE users SET card_id = ? WHERE id = ?"
)

type createChart struct {
}

func MakePostChartClientClient() createChart {
	return createChart{}
}

func (info *createChart) Request(id,face_image, status_image,userID string) error {

	stmt, err := Conn.Prepare(InsertcardInfoQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, face_image, status_image)
	return err

}

func (info *createChart) UpdateRequest(id,userID string) error {

	stmt, err := Conn.Prepare(UpdatUserInfo)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id,userID)
	return err
}