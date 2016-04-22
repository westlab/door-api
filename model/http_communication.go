package model

import (
	"time"
)

// HTTPCommunication is generated HTTP request and response
type HTTPCommunication struct {
	SrcIP   string
	DstIP   string
	SrcPort int64
	DstPort int64
	Host    string
	URI     string
	Title   string
	Time    time.Time
}

// URL return URL generated from Host and URI
func (h *HTTPCommunication) URL() string {
	if h.Host != "" && h.URI != "" {
		return h.Host + h.URI
	}
	return ""
}

func GetIPRank() []Count {
	// pa
}
