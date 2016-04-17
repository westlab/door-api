package model

import (
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
