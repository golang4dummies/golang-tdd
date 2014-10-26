package main

import (
	"fmt"
	"io/ioutil"
	)

func main() {
	fmt.Println("Hello QR Code");	fmt.Println("GenerateQRCode is empty")
	
	qrcode := GenerateQRCode("555-2368")
	ioutil.WriteFile("qrcode.png", qrcode, 0644)
}

func GenerateQRCode(code string) []byte {
	return nil
}
