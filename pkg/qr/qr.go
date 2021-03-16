package qr

import (
	"encoding/base64"
	"fmt"
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"image/png"
	"log"
	"os"
)

func CreateQRCode(userID string)string{
	qrCode, err := qr.Encode(userID, qr.M, qr.Auto)
	if err != nil {
		log.Println(err)
	}

	qrCode, err = barcode.Scale(qrCode, 200, 200)
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Create("qr_"+userID+".png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	err = png.Encode(file, qrCode)
	if err != nil {
		log.Fatal(err)
	}
	///////////////base64へ
	file, _ = os.Open("qr_"+userID+".png")
	defer file.Close()

	fi, _ := file.Stat() //FileInfo interface
	size := fi.Size()    //ファイルサイズ

	data := make([]byte, size)
	file.Read(data)
	if err := os.Remove("qr_"+userID+".png"); err != nil {
		fmt.Println(err)
	}
	return base64.StdEncoding.EncodeToString(data)
}


