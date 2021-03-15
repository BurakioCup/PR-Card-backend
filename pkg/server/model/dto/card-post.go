package dto

import "encoding/base64"


type Chart struct{
	ItemName  []string      `json:"itemName"`
	ItemScore []int     `json:"itemScore"`
}

type RequestCardOver struct {
	FaceImage base64.Encoding        `json:"faceImage"`
    Status    []Chart                `json:"status"`
}
