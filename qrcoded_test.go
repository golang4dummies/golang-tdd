package main //golang-tdd

import (
	"testing"
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
