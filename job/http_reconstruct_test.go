package job

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConverStrToDPI(t *testing.T) {
	data := "1.1.1.1,2.2.2.2,80,12345,2016-04-01 00:00:00,501, /home HTTP 1.1"
	dpi, _ := convertStrToDPI(&data)

	assert.Equal(t, "1.1.1.1", dpi.SrcIP)
	assert.Equal(t, "2.2.2.2", dpi.DstIP)
	assert.Equal(t, int64(80), dpi.SrcPort)
	assert.Equal(t, int64(12345), dpi.DstPort)
	assert.Equal(t, "GET", dpi.Rule)
	assert.Equal(t, " /home HTTP 1.1", dpi.Data)
}
