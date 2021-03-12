package view

type ReadQRResponse struct{
	CardImage string `json:"qrImage"`
}

func ReturnReadQRResponse(imagePath string) ReadQRResponse {
	body := ReadQRResponse{
		CardImage: imagePath,
	}
	return body
}