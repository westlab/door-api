package model

import (
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
	RuleID    int64     `db:"rule_id" json:"rule_id"`
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
