package model

import (
	"testing"
	"time"
)

func TestURL(t *testing.T) {
	hc := HTTPCommunication{
		SrcIP: "1.1.1.1", DstIP: "1.1.1.2",
		SrcPort: 1234, DstPort: 80,
		Host: "google.com", URI: "/foo", Time: time.Now(),
	}
	actual := hc.URL()
	expected := "google.com/foo"
	if actual != expected {
		t.Errorf("got %v\n want %v", actual, expected)
	}
}
