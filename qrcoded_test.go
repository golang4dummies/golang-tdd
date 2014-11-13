package main //golang-tdd

import (
	"testing"
	"bytes"
	"image/png"
)

func TestGenerateQRCodeReturnsValue(t *testing.T) {
	result := GenerateQRCode("555-2368")

	if result == nil {
		t.Errorf("Generated QRCode is nil")
	}

	if len(result) == 0 {
		t.Errorf("Generated QR code lenght is zero")
	}
}

func TestGenerateQRCodeGeneratesPNG(t *testing.T) {
	result := GenerateQRCode("555-2368")
	buffer := bytes.NewBuffer(result)
	_, err := png.Decode(buffer)
	
	if err != nil {
		t.Errorf("Generated QR Code is not a PNG: %s", err)
	}
}
