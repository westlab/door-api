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
		StreamID: 1, Rule: "bar", Timestamp: time.Now(),
		Data: "data",
	}

	assert.Equal(t, "1.1.1.12.2.2.28012345", d.ToKey())
}

func TestParseGET(t *testing.T) {
	data := " /foo/bar/baz HTTP 1.1"
	r, err := parseGET(data)
	assert.Equal(t, nil, err)
	assert.Equal(t, "/foo/bar/baz", r)
	// TODO: Blacktest
}

func TestParseHOST(t *testing.T) {
	data := " google.com \r\nFrom: user@example.com"
	r, err := parseHOST(data)
	assert.Equal(t, nil, err)
	assert.Equal(t, "google.com", r)
	// TODO: Blacktest
}

func TestParseContestType(t *testing.T) {
	data := " application/json \r\nFrom: user@example.com"
	r, err := parseHOST(data)
	assert.Equal(t, nil, err)
	assert.Equal(t, "application/json", r)
	// TODO: Blacktest
}

func TestParseTitle(t *testing.T) {
	data := "> westlab </title><foo>bar</foo>"
	r, err := parseTitle(data)
	assert.Equal(t, nil, err)
	assert.Equal(t, "westlab", r)

	data = "> westlab"
	r, err = parseTitle(data)
	assert.Equal(t, nil, err)
	assert.Equal(t, "westlab", r)
	// TODO: Blacktest
}

func TestParseData(t *testing.T) {
	d := DPI{
		ID: 1, SrcIP: "1.1.1.1", DstIP: "2.2.2.2",
		SrcMac: "abcd", DstMac: "foo", SrcPort: 80, DstPort: 12345,
		StreamID: 1, Rule: "GET", Timestamp: time.Now(),
		Data: " /foo/bar/baz HTTP 1.1\r\n",
	}
	r, err := d.ParseData()
	assert.Equal(t, nil, err)
	assert.Equal(t, "/foo/bar/baz", r)

	d.Rule = "Host:"
	d.Data = " google.com\r\nFrom: user@example.com"
	r, err = d.ParseData()
	assert.Equal(t, nil, err)
	assert.Equal(t, "google.com", r)

	d.Rule = "Content-Type:"
	d.Data = " application/json \r\nFrom: user@example.com"
	r, err = d.ParseData()
	assert.Equal(t, nil, err)
	assert.Equal(t, "application/json", r)

	d.Rule = "<title"
	d.Data = "> westlab </title><foo>bar</foo>"
	r, err = d.ParseData()
	assert.Equal(t, nil, err)
	assert.Equal(t, "westlab", r)
}

func TestToHTTPCommunication(t *testing.T) {
	d := DPI{
		ID: 1, SrcIP: "1.1.1.1", DstIP: "2.2.2.2",
		SrcMac: "abcd", DstMac: "foo", SrcPort: 80, DstPort: 12345,
		StreamID: 1, Rule: "GET", Timestamp: time.Now(),
		Data: " /foo/bar/baz HTTP 1.1\r\n",
	}

	tc := d.ToHTTPCommunication()
	assert.Equal(t, tc.URI, "/foo/bar/baz")
	assert.Equal(t, tc.SrcIP, d.SrcIP)
	assert.Equal(t, tc.SrcPort, d.SrcPort)
	assert.Equal(t, tc.DstIP, d.DstIP)
	assert.Equal(t, tc.DstPort, d.DstPort)

	d.Rule = "Host:"
	d.Data = " google.com\r\nFrom: user@example.com"
	tc = d.ToHTTPCommunication()
	assert.Equal(t, tc.Host, "google.com")

	d.Rule = "Content-Type:"
	d.Data = " application/json \r\nFrom: user@example.com"
	tc = d.ToHTTPCommunication()
	assert.Equal(t, tc.ContentType, "application/json")

	d.Rule = "<title"
	d.Data = "> westlab </title><foo>bar</foo>"
	tc = d.ToHTTPCommunication()
	assert.Equal(t, tc.Title, "westlab")
}
