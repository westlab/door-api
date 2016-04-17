package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestToKey(t *testing.T) {
	d := DPI{
		ID: 1, SrcIP: "1.1.1.1", DstIP: "2.2.2.2",
		SrcMac: "abcd", DstMac: "foo", SrcPort: 80, DstPort: 12345,
		StreamID: 1, RuleID: 1, Rule: "bar", Timestamp: time.Now(),
		Data: "data",
	}

	assert.Equal(t, "1.1.1.12.2.2.28012345", d.ToKey())
}
