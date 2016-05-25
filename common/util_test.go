package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomStr(t *testing.T) {
	s := RandomStr()
	assert.NotZero(t, s)
}

func TestGetMD5Hash(t *testing.T) {
	h1 := GetMD5Hash("http://www.west.sd.keio.ac.jp")
	h2 := GetMD5Hash("http://www.west.sd.keio.ac.jp")
	h3 := GetMD5Hash("http://www.east.sd.keio.ac.jp")
	assert.Equal(t, h1, h2)
	assert.NotEqual(t, h1, h3)
	assert.Equal(t, len(h1), len(h3))
}

func TestIsFileExist(t *testing.T) {
	ok1 := IsFileExist("util.go", ".")
	assert.Equal(t, true, ok1)
	ok2 := IsFileExist("gogogo.go", ".")
	assert.Equal(t, false, ok2)
}
