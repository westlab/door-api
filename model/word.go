package model

import (
	"fmt"

	// lib for mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"

	"github.com/westlab/door-api/conf"
)

// Word is a model of word to track frequency
type Word struct {
	ID    int64  `db:id json:"id"`
	Name  string `db:"name" json:"name"`
	Count int64  `db:"count" json:"count"`
}

// NewWord creates a new Word
func NewWord(name string, count int64) Word {
	return Word{}
}

// Save persists Word in storage
func (w *Word) Save() {
	// TODO: Test case
	// TODO: Share connection among model
	conn, _ := dbr.Open(conf.GetDBType(), conf.GetDSN(), nil)
	// Create session
	sess := conn.NewSession(nil)
	sess.InsertInto("word").Columns("name", "count").
		Record(&w).
		Exec()
}

// GetWordCount returns Word which contains word and count
func GetWordCount(size int64) []Count {
	// TODO: Test case
	var counts []Count

	if size == 0 {
		size = 10000
	}

	// TODO: Share connection among model
	// TODO: Error handling
	conn, _ := dbr.Open(conf.GetDBType(), conf.GetDSN(), nil)
	sess := conn.NewSession(nil)
	sql := fmt.Sprintf(`
		SELECT W.name AS name, SUM(W.count) AS count
		FROM (
		  SELECT name, count FROM word LIMIT %d
		  ) AS W
		GROUP BY name
		ORDER BY count DESC;
	`, size)
	sess.SelectBySql(sql).Load(&counts)
	return counts
}
