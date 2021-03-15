package view

type CreateCardResponse struct {
	Id          string `json:"id"`
	UserName    string `json:"userName"`
	NameImage   string `json:"nameImage"`
	TagImage    string `json:"tagImage"`
	FaceImage   string `json:"faceImage"`
	StatusImage string `json:"statusImage"`
	FreeImage   string `json:"freeImage"`
}


type CreateCardOverResponse struct {
	//Id          string `json:"id"`
	FaceImage   string `json:"faceImage"`
	StatusImage string `json:"statusImage"`
}



func ReturnCreateCardResponse(faceImage, statusImage string)CreateCardOverResponse {
	return CreateCardOverResponse{
		//Id: id,
		FaceImage: faceImage,
		StatusImage: statusImage,
	}

}
