package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomStr(t *testing.T) {
	s := RandomStr()
	assert.NotZero(t, s)
}
