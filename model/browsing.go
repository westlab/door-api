package model

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	// lib for mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/westlab/door-api/common"
	"github.com/westlab/door-api/context"
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

// NewBrowsing creates a new instance of Browsing
func NewBrowsing(SrcIP string, DstIP string, SrcPort int64, DstPort int64,
	Download int64, BrowsingTime int64, Title string, URL string, Domain string,
	Timestamp time.Time) *Browsing {
	nullTime := dbr.NullTime{Time: Timestamp}
	b := Browsing{SrcIP: SrcIP, DstIP: DstIP, SrcPort: SrcPort, DstPort: DstPort,
		Download: Download, BrowsingTime: BrowsingTime, Title: Title, URL: URL,
		Domain: Domain, Timestamp: nullTime}

	return &b
}

// Save saves browsing into db
func (b *Browsing) Save() {
	cxt := context.GetContext()
	conf := cxt.GetConf()
	conn, _ := dbr.Open(conf.DBType, conf.GetDSN(), nil)
	sess := conn.NewSession(nil)
	_, err := sess.InsertInto("browsing").
		Columns("src_ip", "dst_ip", "src_port", "dst_port",
			"timestamp", "download", "browsing_time", "title", "url", "domain").
		Record(*b).
		Exec()
	if err != nil {
		log.Fatal(err)
	}
}

// Update updates browsing in db
func (b *Browsing) Update() (result sql.Result, err error) {
	cxt := context.GetContext()
	conf := cxt.GetConf()
	conn, _ := dbr.Open(conf.DBType, conf.GetDSN(), nil)
	sess := conn.NewSession(nil)

	return sess.Update("browsing").
		Set("src_ip", b.SrcIP).
		Set("dst_ip", b.DstIP).
		Set("src_port", b.SrcPort).
		Set("dst_port", b.DstPort).
		Set("timestamp", b.Timestamp).
		Set("download", b.Download).
		Set("browsing_time", b.BrowsingTime).
		Set("title", b.Title).
		Set("domain", b.Domain).
		Set("url", b.URL).
		Where("id = ?", b.ID).
		Exec()
}

// Delete deletes browsing from db
func (b *Browsing) Delete() {
	// TODO: write sql to delete browsing
}

// MarshalJSON override MarshalJSON for formating json
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

// GetBrowsingBySrcIP returns browsings filtered by src ip
func GetBrowsingBySrcIP(srcIP string) []Browsing {
	var browsings []Browsing
	cxt := context.GetContext()
	conf := cxt.GetConf()
	conn, _ := dbr.Open(conf.DBType, conf.GetDSN(), nil)
	sess := conn.NewSession(nil)

	sess.Select("*").From("browsing").
		Where(
			dbr.And(
				dbr.Eq("src_ip", srcIP),
				dbr.Neq("browsing_time", nil),
			)).
		OrderDir("timestamp", true).
		Load(&browsings)
	return browsings
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
	cxt := context.GetContext()
	conf := cxt.GetConf()

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
	sess.SelectBySql(sql).Load(&browsings)
	return browsings
}

// GetBrowsingHistogram return histogram from browsing
// duration is span from now [minute]
// window is window [minite]
func GetBrowsingHistogram(duraiton int64, window int64) []Count {
	var counts []Count
	view := "browsing_histogram"
	start := time.Now().Add(-time.Duration(duraiton) * time.Minute)
	windows := makeHistogramWindow(start, time.Now(), int64(window))
	countCase := makeCountCase(windows)

	cxt := context.GetContext()
	conf := cxt.GetConf()
	conn, _ := dbr.Open(conf.DBType, conf.GetDSN(), nil)
	sess := conn.NewSession(nil)

	// view is not necessary because we can get value and know type and column name
	// of result. We can create go object in clien side
	// But we want to use something new nad fun.
	// create or replace view
	viewSQL := fmt.Sprintf(`
	    CREATE OR REPLACE VIEW %s AS
		SELECT
		%s
		FROM (SELECT timestamp FROM browsing WHERE timestamp > '%s') AS timestamp_groups
	`, view, countCase, start.Round(10*time.Minute).Format("2006-01-02 15:04:05"))
	sess.Exec(viewSQL)

	unpivots := unpivotHistogram(windows, view)
	selectSQL := fmt.Sprintf(`
		SELECT name, count FROM (
			%s
		) as unpivot_histogram
	`, unpivots)
	sess.SelectBySql(selectSQL).Load(&counts)
	return counts
}

// makeCountCase
// sql formatter
func makeCountCase(windows []common.TimeTuple) string {
	countTemp := "COUNT(CASE WHEN timestamp >= '%s' AND timestamp < '%s' THEN 1 END) AS '%s'"
	num := len(windows)
	countConditions := make([]string, num, num)
	for i, timeTuple := range windows {
		label := fmt.Sprintf("%d:%02d", timeTuple.X.Hour(), timeTuple.X.Minute())
		countConditions[i] = fmt.Sprintf(
			countTemp,
			timeTuple.X.Format("2006-01-02 15:04:05"),
			timeTuple.Y.Format("2006-01-02 15:04:05"),
			label)
	}
	return strings.Join(countConditions, ",\n")
}

// histogramWindow returns window of histogram
func makeHistogramWindow(start time.Time, end time.Time, window int64) []common.TimeTuple {
	start = start.Round(10 * time.Minute)
	end = end.Round(10 * time.Minute)
	span := end.Sub(start)
	windowMinute := time.Duration(window) * time.Minute
	num := int64(span / windowMinute)
	histogramWindow := make([]time.Time, num, num)

	for i := 0; i < int(num); i++ {
		histogramWindow[i] = start.Add(windowMinute * time.Duration(i))
	}

	return common.PairwiseTime(histogramWindow)
}

// unpivotHistogram unpivot table or view for count
func unpivotHistogram(windows []common.TimeTuple, table string) string {
	unpivotTemp := "SELECT '%s' AS name, `%s`.`%s` FROM `%s`"
	num := len(windows)
	unpivots := make([]string, num, num)

	for i, timeTuple := range windows {
		label := fmt.Sprintf("%d:%02d", timeTuple.X.Hour(), timeTuple.X.Minute())
		unpivots[i] = fmt.Sprintf(unpivotTemp, label, table, label, table)
	}

	return strings.Join(unpivots, "\nUNION ALL\n")
}

// GetBrowsingRank retuns count of given column
func GetBrowsingRank(column string, duration int64) []Count {
	var counts []Count
	cxt := context.GetContext()
	conf := cxt.GetConf()
	conn, _ := dbr.Open(conf.DBType, conf.GetDSN(), nil)
	sess := conn.NewSession(nil)
	from := time.Now().Add(-time.Duration(duration))
	sql := fmt.Sprintf(`
		SELECT B.name AS name, SUM(B.%s) AS count
		FROM browsing AS B
		GROUP BY %s
		ORDER BY count DESC
		WHERE B.timestamp >= %s
		`, column, column, from.Format("2015-04-01 11:24:00"))
	sess.SelectBySql(sql).Load(&counts)
	return counts
}

// GetBrowsingAfterID returns Browsing which has id greater than given id
func GetBrowsingAfterID(id int64, limit int64, hasBrowsingTime bool) []Browsing {
	var browsings []Browsing
	var condition dbr.Condition
	cxt := context.GetContext()
	conf := cxt.GetConf()
	conn, _ := dbr.Open(conf.DBType, conf.GetDSN(), nil)
	sess := conn.NewSession(nil)
	if hasBrowsingTime {
		condition = dbr.Gte("id", id)
	} else {
		condition = dbr.And(
			dbr.Gte("id", id),
			dbr.Neq("browsing_time", nil))
	}

	query := sess.Select("*").From("browsing").
		Where(condition).
		OrderDir("timestamp", true)

	if limit > 0 {
		query.Limit(uint64(limit))
	}

	query.Load(&browsings)

	return browsings
}
