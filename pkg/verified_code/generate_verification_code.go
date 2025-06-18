package verified_code

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

func GenerateVerificationCode() (string, error) {
	vcode := make([]byte, 6)
	if _, err := rand.Read(vcode); err != nil {
		log.Printf("Error generating verification code: %v", err)
		return "", err
	}
	return base64.StdEncoding.EncodeToString(vcode), nil
}
