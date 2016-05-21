package model

import (
	"database/sql"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/gchaincl/dotsql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
	"github.com/westlab/door-api/conf"
	"github.com/westlab/door-api/context"
)

type BrowsingTestSuite struct {
	suite.Suite
}

func TestBrowsingTestSuite(t *testing.T) {
	suite.Run(t, new(BrowsingTestSuite))
}

func (s *BrowsingTestSuite) SetupTest() {
	// load conf file
	conf := conf.New("../config.toml")
	context.NewContext(conf)
	cxt := context.GetContext()

	db, err := sql.Open(cxt.GetConf().DBType, cxt.GetConf().GetDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dsql, err := dotsql.LoadFromFile("../test/test_data.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = dsql.Exec(db, "drop-browsing-table")
	if err != nil {
		log.Fatal(err)
	}

	_, err = dsql.Exec(db, "create-browsing-table")
	if err != nil {
		log.Fatal(err)
	}

	_, err = dsql.Exec(db, "insert-browsing-rows")
	if err != nil {
		log.Fatal(err)
	}
}

func (s *BrowsingTestSuite) TearDonwTest() {
}

func (s *BrowsingTestSuite) TestBrowsingCRUD() {
	// Save
	b := NewBrowsing("10.24.1.20", "6.6.6.6", 123, 80, 3, 20, "Mitsubishi", "http://mitsubishi.co.jp", "mitsubishi.co.jp", time.Now())
	b.Save()
	bs := GetBrowsingBySrcIP("10.24.1.20")
	s.Equal("Mitsubishi", bs[0].Title)

	// Update
	b = GetBrowsingByID(int64(2))
	b.Title = "KiraYoshikage"
	_, err := b.Update()
	if err != nil {
		log.Println(err)
	}
	b = GetBrowsingByID(int64(2))
	s.Equal("KiraYoshikage", b.Title)

	// Delete
	// TODO: add delete test
}

func (s *BrowsingTestSuite) TestBrowsingJSON() {
}

func (s *BrowsingTestSuite) TestGetBrowsingByID() {
}

func (s *BrowsingTestSuite) TestGetBrowsingBySrcIP() {
}

func (s *BrowsingTestSuite) TestGetBrowsings() {
}

func (s *BrowsingTestSuite) TestGetBrowsingHistogram() {
}

func (s *BrowsingTestSuite) TestMakeCountCase() {
	start := time.Date(1990, time.November, 24, 0, 0, 0, 0, time.UTC)
	end := start.Add(time.Duration(100) * time.Minute)
	windows := makeHistogramWindow(start, end, int64(10))
	c := makeCountCase(windows)

	expected := `COUNT(CASE WHEN timestamp >= '1990-11-24 00:00:00' AND timestamp < '1990-11-24 00:10:00' THEN 1 END) AS '0:00',
COUNT(CASE WHEN timestamp >= '1990-11-24 00:10:00' AND timestamp < '1990-11-24 00:20:00' THEN 1 END) AS '0:10',
COUNT(CASE WHEN timestamp >= '1990-11-24 00:20:00' AND timestamp < '1990-11-24 00:30:00' THEN 1 END) AS '0:20',
COUNT(CASE WHEN timestamp >= '1990-11-24 00:30:00' AND timestamp < '1990-11-24 00:40:00' THEN 1 END) AS '0:30',
COUNT(CASE WHEN timestamp >= '1990-11-24 00:40:00' AND timestamp < '1990-11-24 00:50:00' THEN 1 END) AS '0:40',
COUNT(CASE WHEN timestamp >= '1990-11-24 00:50:00' AND timestamp < '1990-11-24 01:00:00' THEN 1 END) AS '0:50',
COUNT(CASE WHEN timestamp >= '1990-11-24 01:00:00' AND timestamp < '1990-11-24 01:10:00' THEN 1 END) AS '1:00',
COUNT(CASE WHEN timestamp >= '1990-11-24 01:10:00' AND timestamp < '1990-11-24 01:20:00' THEN 1 END) AS '1:10',
COUNT(CASE WHEN timestamp >= '1990-11-24 01:20:00' AND timestamp < '1990-11-24 01:30:00' THEN 1 END) AS '1:20'`
	s.Equal(expected, c)
}

func (s *BrowsingTestSuite) TestMakeHistogramWindow() {
	start := time.Date(1990, time.November, 24, 0, 0, 0, 0, time.UTC)
	end := start.Add(time.Duration(100) * time.Minute)
	windows := makeHistogramWindow(start, end, int64(10))
	s.Equal(9, len(windows))
	s.Equal(windows[0].Y, windows[1].X)
	s.Equal(windows[7].Y, windows[8].X)
}

func (s *BrowsingTestSuite) TestMakeUnpivotHistogram() {
	start := time.Date(1990, time.November, 24, 0, 0, 0, 0, time.UTC)
	end := start.Add(time.Duration(100) * time.Minute)
	windows := makeHistogramWindow(start, end, int64(10))
	unp := unpivotHistogram(windows, "browsing_histogram")
	unions := []string{
		"SELECT '0:00' AS name, `browsing_histogram`.`0:00` FROM `browsing_histogram`",
		"UNION ALL",
		"SELECT '0:10' AS name, `browsing_histogram`.`0:10` FROM `browsing_histogram`",
		"UNION ALL",
		"SELECT '0:20' AS name, `browsing_histogram`.`0:20` FROM `browsing_histogram`",
		"UNION ALL",
		"SELECT '0:30' AS name, `browsing_histogram`.`0:30` FROM `browsing_histogram`",
		"UNION ALL",
		"SELECT '0:40' AS name, `browsing_histogram`.`0:40` FROM `browsing_histogram`",
		"UNION ALL",
		"SELECT '0:50' AS name, `browsing_histogram`.`0:50` FROM `browsing_histogram`",
		"UNION ALL",
		"SELECT '1:00' AS name, `browsing_histogram`.`1:00` FROM `browsing_histogram`",
		"UNION ALL",
		"SELECT '1:10' AS name, `browsing_histogram`.`1:10` FROM `browsing_histogram`",
		"UNION ALL",
		"SELECT '1:20' AS name, `browsing_histogram`.`1:20` FROM `browsing_histogram`"}
	expected := strings.Join(unions, "\n")
	s.Equal(expected, unp)
}

func (s *BrowsingTestSuite) TestGetBrowsingRank() {
}

func (s *BrowsingTestSuite) TestGetBrowsingAfterID() {
}
