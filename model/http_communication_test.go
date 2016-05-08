package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	hc := HTTPCommunication{
		SrcIP: "1.1.1.1", DstIP: "1.1.1.2",
		SrcPort: 1234, DstPort: 80, Title: "foo",
		Host: "google.com", URI: "/foo", Time: time.Now(),
	}
	actual := hc.URL()
	expected := "google.com/foo"
	if actual != expected {
		t.Errorf("got %v\n want %v", actual, expected)
	}
}

func TestIsValid(t *testing.T) {
	hc := HTTPCommunication{
		SrcIP: "", DstIP: "1.1.1.2",
		SrcPort: 1234, DstPort: 80, Title: "foo",
		Host: "google.com", URI: "/foo", Time: time.Now(),
	}
	assert.Equal(t, false, hc.IsValid(""))
	hc.SrcIP = "1.1.1.1"
	assert.Equal(t, false, hc.IsValid(""))
	hc.ContentType = "text/html"
	assert.Equal(t, true, hc.IsValid(""))
}

func TestToBrowsing(t *testing.T) {
	hc := HTTPCommunication{
		SrcIP: "1.1.1.1", DstIP: "1.1.1.2",
		SrcPort: 1234, DstPort: 80, Title: "foo",
		Host: "google.com", URI: "/foo", Time: time.Now(),
	}
	b := hc.ToBrowsing()

	assert.Equal(t, hc.SrcIP, b.SrcIP)
	assert.Equal(t, hc.DstIP, b.DstIP)
	assert.Equal(t, hc.SrcPort, b.SrcPort)
	assert.Equal(t, hc.DstPort, b.DstPort)
	assert.Equal(t, hc.URL(), b.URL)
	assert.Equal(t, hc.Title, b.Title)
}
