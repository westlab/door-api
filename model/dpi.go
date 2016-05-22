package model

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// DPI is a model for storing Deep Packet Inspection Result
type DPI struct {
	ID        int64     `db:"id" json:"id"`
	SrcIP     string    `db:"src_ip" json:"src_ip"`
	DstIP     string    `db:"dst_ip" json:"dst_ip"`
	SrcMac    string    `db:"src_mac" json:"src_mac"`
	DstMac    string    `db:"dst_mac" json:"dst_mac"`
	SrcPort   int64     `db:"src_port" json:"src_port"`
	DstPort   int64     `db:"dst_port" json:"dst_port"`
	StreamID  int64     `db:"stream_id" json:"stream_id"`
	Rule      string    `db:"rule" json:"rule"`
	Timestamp time.Time `db:"timestamp" json:"timestamp"`
	Data      string    `db:"data" json:"data"`
}

// ToKey returns key generated from ips and ports
func (d *DPI) ToKey() string {
	var portStr string
	var ipStr string
	if d.SrcPort <= d.DstPort {
		portStr = strconv.Itoa(int(d.SrcPort)) + strconv.Itoa(int(d.DstPort))
	} else {
		portStr = strconv.Itoa(int(d.DstPort)) + strconv.Itoa(int(d.SrcPort))
	}
	if strings.Compare(d.SrcIP, d.DstIP) == -1 {
		ipStr = d.SrcIP + d.DstIP
	} else {
		ipStr = d.DstIP + d.SrcIP
	}
	return ipStr + portStr
}

// ToHTTPCommunication converts DPI to HTTPCommunication
func (d *DPI) ToHTTPCommunication() *HTTPCommunication {
	tc := HTTPCommunication{SrcIP: d.SrcIP, DstIP: d.DstIP,
		SrcPort: d.SrcPort, DstPort: d.DstPort, Time: time.Now()}
	data, _ := d.ParseData()
	switch d.Rule {
	case "GET":
		tc.URI = data
	case "Host:":
		tc.Host = data
	case "Content-Type:":
		tc.ContentType = data
	case "<title":
		tc.Title = data
	}
	return &tc
}

// ParseData parses data according to the rule
// Return Error when fail to parse
//
// rules are supposed to follows:
//   GET
//   HOST:
//   Content-Type:
//   <title
func (d *DPI) ParseData() (string, error) {
	switch d.Rule {
	case "GET":
		return parseGET(d.Data)
	case "Host:":
		return parseHOST(d.Data)
	case "Content-Type:":
		return parseContentType(d.Data)
	case "<title":
		return parseTitle(d.Data)
	default:
		return d.Data, errors.New("Rule does not match")
	}
}

// parseGET parses HTTP GET request and extracts URI
func parseGET(s string) (string, error) {
	get := strings.SplitN(s, "\r\n", 2)
	uriHTTP := strings.TrimSpace(get[0])
	idx := strings.Index(uriHTTP, " ")
	if idx == -1 {
		return uriHTTP, errors.New("URI is too long")
	}
	return uriHTTP[:idx], nil
}

// parseHOST parses HTTP Header and extracts HOST
func parseHOST(s string) (string, error) {
	host := strings.SplitN(s, "\r\n", 2)
	return strings.TrimSpace(host[0]), nil
}

// parseContentType parses HTTP header and extracts Content-Type
//
// data candidate
// text/html
// text/*
// text: text/html; charset=ISO-8859-1
func parseContentType(s string) (string, error) {
	contentTypeRaw := strings.SplitN(s, "\r\n", 2)
	contentType := contentTypeRaw[0]
	idx := strings.Index(contentType, ";")
	if idx == -1 {
		return strings.TrimSpace(contentType), nil
	}
	return strings.TrimSpace(contentType[:idx]), nil
}

// parseTitle parses HTTP body and extract title
func parseTitle(s string) (string, error) {
	first := strings.Index(s, ">")
	if first == -1 {
		return "", errors.New("title may be too long or failed to capture data")
	}
	end := strings.Index(s, "</")
	if end == -1 {
		return strings.TrimSpace(s[first+1:]), nil
	}
	return strings.TrimSpace(s[first+1 : end]), nil
}
