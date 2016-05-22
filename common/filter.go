package common

import (
	"io/ioutil"
	"log"
	"strings"
)

const (
	zeroByte uint8 = 0
)

// Filter for restrict banned word
type BlackListFilter struct {
	blackList []string
}

// NewFilter creates Filter instance
func NewBlackListFilter(blackListFile string) *Filter {
	// If file is not too large, ReadFile is better solution
	// If not, use buffer io
	file, err := ioutil.ReadFile(blackListFile)
	if err != nil {
		log.Println(err)
	}

	f := func(c rune) bool {
		return c == '\n'
	}

	blackList := strings.FieldsFunc(string(file), f)
	return &Filter{blackList}
}

// Ok returns true if word is not matched in blackList
func (f *Filter) Ok(word string) (ok bool) {
	for _, b := range f.blackList {
		if strings.Contains(word, b) {
			return false
		}
	}
	return true
}

// IsZeroByte check if string is zeroByte
func IsZeroByte(s string) bool {
	if zeroByte == uint8(s) {
		return true
	}
	return false
}
