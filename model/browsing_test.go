package model

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMakeCountCase(t *testing.T) {
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
	assert.Equal(t, expected, c)
}

func TestMakeHistogramWindow(t *testing.T) {
	start := time.Date(1990, time.November, 24, 0, 0, 0, 0, time.UTC)
	end := start.Add(time.Duration(100) * time.Minute)
	windows := makeHistogramWindow(start, end, int64(10))
	assert.Equal(t, 9, len(windows))
	assert.Equal(t, windows[0].Y, windows[1].X)
	assert.Equal(t, windows[7].Y, windows[8].X)
}

func TestMakeUnpivotHistogram(t *testing.T) {
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
	assert.Equal(t, expected, unp)
}