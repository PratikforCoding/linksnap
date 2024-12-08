package utils

import (
	"encoding/base64"

	"github.com/skip2/go-qrcode"
)

// GenerateQRCode generates a QR code for a given URL
func GenerateQRCode(url string) (string, error) {
	var png []byte
	png, err := qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}

	base64Image := base64.StdEncoding.EncodeToString(png)
	return "data:image/png;base64," + base64Image, nil
}