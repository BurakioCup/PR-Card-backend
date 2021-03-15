package dto

type Chart struct {
	ItemName  [5]string `json:"itemName"`
	ItemScore [5]int    `json:"itemScore"`
}

type RequestCardOver struct {
	FaceImage string          `json:"faceImage"`
	Status    Chart         `json:"status"`
}

//takashiへのリクエストボディ
type RequestCardOverNode struct {
	FaceImage string        `json:"faceImage"`
	FaceImageName string     `json:"faceIconName"`
	Status     Chart         `json:"status"`
	StatusImageName string   `json:"chartName"`
}

func RequestCardResponse(cardID, faceImage string, status Chart) RequestCardOverNode {
	return RequestCardOverNode{
		FaceImage:       faceImage,
		FaceImageName:   cardID + "_" + "faceImage" + ".png",
		Status:          status,
		StatusImageName: cardID + "_" + "statusImage" + ".png",
	}
}






