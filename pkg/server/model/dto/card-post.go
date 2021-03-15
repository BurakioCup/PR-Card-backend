package dto

type Chart struct {
	ItemName  []string `json:"itemName"`
	ItemScore []int    `json:"itemScore"`
}

type RequestCardOver struct {
	FaceImage string        `json:"faceImage"`
	Status    Chart         `json:"status"`
}

type RequestCardOverSample struct {
	FaceImage string `json:"faceImage"`
	ItemName  []string `json:"itemName"`
	ItemScore []int    `json:"itemScore"`
}

//takashiへのリクエストボディ
type RequestCardOverNode struct {
	FaceImage string `json:"faceImage"`
	FaceImageName string      `json:"faceIconName"`
	Status    Chart         `json:"status"`
	StatusImageName string    `json:"chartName"`
}


type CreateCardOverResponse struct {
	//Id          string `json:"id"`
	FaceImage   string `json:"faceImage"`
	StatusImage string `json:"statusImage"`
}

/*
func (r RequestCardOverNode) Read(p []byte) (n int, err error) {
	panic("implement me")
}

 */

func RequestCardResponse(cardID,faceImage string,status Chart)RequestCardOverNode{
	return RequestCardOverNode{
		FaceImage :faceImage,
		FaceImageName : cardID+"_"+"faceImage"+".png",
		Status  : status,
		StatusImageName:cardID+"_"+"statusImage"+".png",
	}
}






func (r RequestCardOver) Read(p []byte) (n int, err error) {
	panic("implement me")
}



