package dto

type RequestItemScore struct {
	ItemName  string `json:"itemName"`
	ItemScore string `json:"itemScore"`
}

type RequestCardOver struct {
	UserName  string             `json:"userName"`
	FaceImage string             `json:"faceImage"`
	Status    []RequestItemScore `json:"status"`
}

