package main

import (
	"fmt"
	"io/ioutil"
	"image"
	"image/png"
	"bytes"
)

func main() {
	fmt.Println("Hello QR Code")

	qrcode := GenerateQRCode("555-2368")
	fmt.Printf("GenerateQRCode is %v", qrcode)
	ioutil.WriteFile("qrcode.png", qrcode, 0644)
}

func GenerateQRCode(code string) []byte {
	img := image.NewNRGBA(image.Rect(0,0, 21,21))
	buf := new(bytes.Buffer)
	_ = png.Encode(buf, img)
	
	return buf.Bytes()
}
