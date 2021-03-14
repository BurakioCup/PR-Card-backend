package dao



const (
	InsertcardWordInfoQuery = "INSERT INTO `card_my_word` (id,card_id,word) VALUES (?,?,?)"
)


type createWord struct {
}

func MakePostWordClientClient() createWord {
	return createWord{}
}

func (info *createWord) Request(id,word string) error {
//仮データ
	card_id:="001"
	stmt, err := Conn.Prepare(InsertcardWordInfoQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id,card_id,word)
	return err
}