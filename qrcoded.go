package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello QR Code")

	qrcode := GenerateQRCode("555-2368")
	fmt.Printf("GenerateQRCode is %v", qrcode)
	ioutil.WriteFile("qrcode.png", qrcode, 0644)
}

func GenerateQRCode(code string) []byte {
	return []byte{0xFF}
}
