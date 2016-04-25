package model

import (
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

// New creates a new instance of Browsing
func New(SrcIP string, DstIP string, SrcPort int64, DstPort int64,
	Download int64, BrowsingTime int64, Title string, URL string) Browsing {

	return Browsing{}
}

// Save saves browsing into db
func (b *Browsing) Save() {
	// TODO: write sql to save browsing to db
}

// Update updates browsing in db
func (b *Browsing) Update() {
	// TODO: write sql to update browsing
}

// Delete deletes browsing from db
func (b *Browsing) Delete() {
	// TODO: write sql to delete browsing
}

// GetBrowsingByID returns Browsing by id
func GetBrowsingByID(id int64) Browsing {
	// TODO: write sql code and load it to Browsing
	return Browsing{}
}

// GetBrowsings returns list of Browsing
func GetBrowsings(q string, size int64) []Browsing {
	// TODO: write sql code and load it to the array of Browsing
	// params:
	// q is a search sring and search title, url, src_ip, dst_ip, src_port, dst_port
	// with regular expression
	//
	// size is a number of the result
	var browsings []Browsing
	return browsings
}
