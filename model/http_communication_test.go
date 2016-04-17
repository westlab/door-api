package model

import (
	"testing"
	"time"
)

// var hcs = []HTTPCommunication{
// 	HTTPCommunication{SrcIP: "1.1.1.1", DstIP: "1.1.1.2", SrcPort: 1234, DstPort: 80, Host: "google.com", URI: "/foo", Time: time.Now()},
// 	HTTPCommunication{SrcIP: "1.1.1.2", DstIP: "1.1.1.1", SrcPort: 80, DstPort: 1234, Host: "google.com", URI: "/foo", Time: time.Now()},
// }

func TestURL(t *testing.T) {
	hc := HTTPCommunication{SrcIP: "1.1.1.1", DstIP: "1.1.1.2", SrcPort: 1234, DstPort: 80, Host: "google.com", URI: "/foo", Time: time.Now()}
	actual := hc.URL()
	expected := "1.1.1.11.1.1.2801234"
	if actual != expected {
		t.Errorf("got %v\n want %v")
	}
}
