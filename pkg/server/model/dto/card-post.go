package dto

type RequestItem struct {
	ItemName  string `json:"itemName"`
	ItemScore int    `json:"itemScore"`
}

type RequestCardOver struct {
	UserName  string        `json:"userName"`
	FaceImage string        `json:"faceImage"`
	Status    []RequestItem `json:"status"`
}
