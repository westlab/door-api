package common

import (
	"crypto/rand"
	"encoding/base64"
	"log"
)

// RandomStr generate random string
func RandomStr() string {
	size := 32

	rb := make([]byte, size)
	_, err := rand.Read(rb)

	if err != nil {
		log.Println(err)
	}

	return base64.URLEncoding.EncodeToString(rb)
}
