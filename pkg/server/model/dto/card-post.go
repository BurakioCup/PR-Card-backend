package dto


type RequestWord struct {
	Word  string `json:"word"`
}

type RequestCardDetail struct {
	Nickname  string             `json:"nickName"`
	Words     []RequestWord      `json:"words"`
	freeText    string           `json:"freeText"`
}

