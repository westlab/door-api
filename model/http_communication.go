package model

import (
	"time"
)

// HTTPCommunication is generated HTTP request and response
type HTTPCommunication struct {
	SrcIP       string
	DstIP       string
	SrcPort     int64
	DstPort     int64
	Host        string
	URI         string
	Title       string
	ContentType string
	Time        time.Time
}

// URL return URL generated from Host and URI
func (h *HTTPCommunication) URL() string {
	if h.Host != "" && h.URI != "" {
		return h.Host + h.URI
	}
	return ""
}

// IsValid validate http communication
// TODO: reconsider interface
func (h *HTTPCommunication) IsValid(contentType string) bool {
	if contentType == "" {
		contentType = "text/html"
	}

	if h.SrcIP != "" && h.DstIP != "" &&
		h.SrcPort != 0 && h.DstPort != 0 &&
		h.URL() != "" && h.Title != "" &&
		h.ContentType == contentType {
		return true
	}
	return false
}
