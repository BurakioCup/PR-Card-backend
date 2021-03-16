package dto

type Chart struct {
	ItemName  [5]string `json:"itemName"`
	ItemScore [5]int    `json:"itemScore"`
}

type RequestCardOver struct {
	FaceImage string        `json:"faceImage"`
	Status    Chart         `json:"status"`
}

type RequestCardOverSample struct {
	FaceImage string `json:"faceImage"`
	ItemName  [5]string `json:"itemName"`
	ItemScore [5]int    `json:"itemScore"`
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

func RequestCardResponse(cardID,faceImage string,itemName [5]string,itemScore [5]int)RequestCardOverNode{
	var chart Chart
	for i := 0; i < 5; i++ {
		chart.ItemName[i] = itemName[i]
		chart.ItemScore[i] = itemScore[i]
	}
	return RequestCardOverNode{
		FaceImage :faceImage,
		FaceImageName : cardID+"_"+"faceImage"+".png",
		Status  : chart,
		StatusImageName:cardID+"_"+"statusImage"+".png",
	}
}

func (r RequestCardOver) Read(p []byte) (n int, err error) {
	panic("implement me")
}

type CardDetailRequest struct{
	Name string `json:"name"`
	NickName string `json:"nickName"`
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
	var userName UserName
	userName.Name=req.Name
	userName.NickName=req.NickName
	return CardDetailNode{
		UserName: userName,
		UserNameFileN : cardID+"_"+"nameImage"+".png",
		HashTags: req.HashTags,
		HashTagsFileN:cardID+"_"+"tagImage"+".png",
		FreeText: req.FreeText,
		FreeTextFileN:cardID+"_"+"freeImage"+".png",
	}
}