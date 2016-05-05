package model

import (
	"fmt"

	// lib for mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/westlab/door-api/conf"
)

// Browsing is a model for http browsing
type Browsing struct {
	ID           int64        `db:"id" json:"id"`
	SrcIP        string       `db:"src_ip" json:"src_ip"`
	DstIP        string       `db:"dst_ip" json:"dst_ip"`
	SrcPort      int64        `db:"src_port" json:"src_port"`
	DstPort      int64        `db:"dst_port" json:"dst_port"`
	Timestamp    dbr.NullTime `db:"timestamp" json:"timestamp"`
	CreatedAt    dbr.NullTime `db:"created_at" json:"created_at"`
	Download     int64        `db:"download" json:"download"`
	BrowsingTime int64        `db:"browsing_time" json:"browsing_time"`
	Title        string       `db:"title" json:"title"`
	URL          string       `db:"url" json:"url"`
	Domain       string       `db:"domain" json:"domain"`
}

// New creates a new instance of Browsing
func New(SrcIP string, DstIP string, SrcPort int64, DstPort int64,
	Download int64, BrowsingTime int64, Title string, URL string, Domain string) Browsing {

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

func (b *Browsing) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(
		"{\"id\":\"%d\", \"src_ip\":\"%s\", \"dst_ip\":\"%s\", \"src_port\":\"%d\", \"dst_port\":\"%d\", \"timestamp\":\"%s\", \"created_at\":\"%s\", \"download\":\"%d\", \"browsing_time\":\"%d\", \"title\":\"%s\", \"url\":\"%s\", \"domain\":\"%s\"}",
		b.ID,
		b.SrcIP,
		b.DstIP,
		b.SrcPort,
		b.DstPort,
		b.Timestamp.Time.Format("2006-01-02 15:04:05"),
		b.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		b.Download,
		b.BrowsingTime,
		b.Title,
		b.URL,
		b.Domain)), nil
}

// GetBrowsingByID returns Browsing by id
func GetBrowsingByID(id int64) Browsing {
	// TODO: write sql code and load it to Browsing
	return Browsing{}
}

// GetBrowsings returns list of Browsing
func GetBrowsings(q string, size int64) []Browsing {
	// params:
	// q is a search sring and search title, url, src_ip, dst_ip, src_port, dst_port
	// with regular expression
	//
	// size is a number of the result
	var browsings []Browsing
	var sql string
	conf := conf.GetConf()

	if size == 0 {
		size = 100
	}

	conn, _ := dbr.Open(conf.DBType, conf.GetDSN(), nil)
	sess := conn.NewSession(nil)
	if q != "" {
		sql = fmt.Sprintf(`
			SELECT id, src_ip, dst_ip, src_port, dst_port, timestamp, created_at, download, browsing_time, title, url, domain
			FROM browsing
			WHERE browsing_time IS NOT NULL
				AND (
					src_ip REGEXP '%s'
					OR dst_ip REGEXP '%s'
					OR src_port REGEXP '%s'
					OR dst_port REGEXP '%s'
					OR title REGEXP '%s'
					OR url REGEXP '%s'
				)
			ORDER BY id DESC
			LIMIT %d;
			`, q, q, q, q, q, q,
			size)
	} else {
		sql = fmt.Sprintf(`
			SELECT id, src_ip, dst_ip, src_port, dst_port, timestamp, created_at, download, browsing_time, title, url, domain
			FROM browsing
			WHERE browsing_time IS NOT NULL
			ORDER BY id DESC
			LIMIT %d;
			`, size)
	}
	sess.SelectBySql(sql).LoadStruct(&browsings)
	return browsings
}
