package model

import (
	// lib for mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
	"time"
)

// Browsing is a model for http browsing
type Browsing struct {
	ID           int64     `db:"id" json:"id"`
	SrcIP        string    `db:"src_ip" json:"src_ip"`
	DstIP        string    `db:"dst_ip" json:"dst_ip"`
	SrcPort      int64     `db:"src_port" json:"src_port"`
	DstPort      int64     `db:"dst_port" json:"dst_port"`
	Timestamp    time.Time `db:"timestamp" json:"timestamp"`
	CreatedAt    time.Time `db:"create_at" json:"created_at"`
	Download     int64     `db:"download" json:"download"`
	BrowsingTime int64     `db:"browsing_time" json:"browsing_time"`
	Title        string    `db:"title" json:"title"`
	URL          string    `db:"url" json:"url"`
}
