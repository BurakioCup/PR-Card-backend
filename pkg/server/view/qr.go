package view

type ReadQRResponse struct{
	cardImage string `json:"qrImage"`
}

func ReturnReadQRResponse(imagePath string) ReadQRResponse {
	body := ReadQRResponse{
		cardImage: imagePath,
	}
	return body
}