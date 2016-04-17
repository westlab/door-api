package model

import (
	"strconv"
	"strings"
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
}

// ToKey
// func (h *HTTPCommunication) ToKey() string {
// 	var portStr string
// 	var ipStr string
//
// 	if h.SrcPort <= h.DstPort {
// 		portStr = strconv.Itoa(h.SrcPort) + strconv.Itoa(h.DstPort)
// 	} else {
// 		portStr = strcov.Itoa(h.DstPort) + strconv.Itoa(h.SrcPort)
// 	}
// 	if strings.Compare(h.SrcIP, h.DstPort) == -1 {
// 		ipStr = h.SrcIP + h.DstIP
// 	} else {
// 		ipStr = h.DstIP + h.SrcIP
// 	}
// 	return ipStr + portStr
// }
