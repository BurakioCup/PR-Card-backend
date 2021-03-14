package dao



const (
	InsertcardInfoQuery    = "INSERT INTO `cards` (id,name,image_path) VALUES (?,?,?)"
	InsertcardWordInfoQuery = "INSERT INTO `card_my_word` (id,card_id,word) VALUES (?,?,?)"
)


type createWord struct {
}

func MakePostCardClientClient() createWord {
	return createWord{}
}

func (info *createWord) Request(id,card_id,word string) error {

	stmt, err := Conn.Prepare(InsertcardWordInfoQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id,card_id,word)
	return err
}