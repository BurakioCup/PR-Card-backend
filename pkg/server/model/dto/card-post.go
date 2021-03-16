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

type CardDetailRequest struct{
	UserName UserName `json:"userName"`
	HashTags [4]string `json:"hashTags"`
	FreeText string `json:"freeText"`
}

type UserName struct{
	Name string `json:"name"`
	NickName string `json:"nickName"`
}

type CardDetailNode struct{
	UserName UserName `json:"userName"`
	UserNameFileN string `json:"userNameFileN"`
	HashTags [4]string `json:"hashTags"`
	HashTagsFileN string `json:"hashTagsFileN"`
	FreeText string `json:"freeText"`
	FreeTextFileN string `json:"freeTextFileN"`
}

func RequestCardDetailResponse(cardID string,req CardDetailRequest)CardDetailNode{
	return CardDetailNode{
		UserName: req.UserName,
		UserNameFileN : cardID+"_"+"nameImage"+".png",
		HashTags: req.HashTags,
		HashTagsFileN:cardID+"_"+"tagImage"+".png",
		FreeText: req.FreeText,
		FreeTextFileN:cardID+"_"+"freeImage"+".png",
	}
}