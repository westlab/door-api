package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectSingleMeta(t *testing.T) {
	m := SelectSingleMeta("door-version")
	assert.Equal(t, "door-version", m.Name)
	assert.Equal(t, "1.0", m.Value)
	assert.Equal(t, "2016-04-01 00:00:00", m.CreatedAt.Time.Format("2006-01-02 15:04:05"))
}
