package common

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	blackListFile, _ := ioutil.TempFile(os.TempDir(), "door-api")
	blackList := `
foo
bar
baz
	`
	ioutil.WriteFile(blackListFile.Name(), []byte(blackList), os.ModePerm)

	f := NewBlackListFilter(blackListFile.Name())
	assert.Equal(t, f.Ok("foooo"), false)
	assert.Equal(t, f.Ok("foooo"), false)
	assert.Equal(t, f.Ok("hoge"), true)
}
