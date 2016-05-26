package common

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"log"
	"os"
	"path/filepath"
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

// GetMD5Hash gets hash string from a string
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// IsFileExist checks if the file name is existed
func IsFileExist(fileName string, dirPath string) bool {
	fpath := filepath.Join(dirPath, fileName)
	_, err := os.Stat(fpath)
	return err == nil
}
