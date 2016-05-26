package model

import (
	"fmt"
	"log"

	// lib for mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/westlab/door-api/context"
)

// Word is a model of word to track frequency
type Word struct {
	ID    int64  `db:id json:"id"`
	Name  string `db:"name" json:"name"`
	Count int64  `db:"count" json:"count"`
}

// NewWord creates a new Word
func NewWord(name string, count int64) *Word {
	return &Word{Name: name, Count: count}
}

// Save persists Word in storage
func (w *Word) Save() {
	// TODO: Test case
	// TODO: Share connection among model
	cxt := context.GetContext()

	conn, _ := dbr.Open(cxt.GetConf().DBType, cxt.GetConf().GetDSN(), nil)
	// Create session
	sess := conn.NewSession(nil)
	defer sess.Close()
	_, err := sess.InsertInto("word").Columns("name", "count").
		Record(*w).
		Exec()
	if err != nil {
		log.Println(err)
	}
}

// WordBulkInsert saves words to word table with bulk insert
func WordBulkInsert(words []*Word) error {
	cxt := context.GetContext()

	conn, _ := dbr.Open(cxt.GetConf().DBType, cxt.GetConf().GetDSN(), nil)
	// Create session
	sess := conn.NewSession(nil)
	defer sess.Close()

	stmt := sess.InsertInto("word").Columns("name", "count")
	for _, word := range words {
		stmt.Record(word)
	}

	_, err := stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}

// GetWordCount returns Word which contains word and count
func GetWordCount(size int64) []Count {
	// TODO: Test case
	var counts []Count
	cxt := context.GetContext()

	if size == 0 {
		size = 10000
	}

	// TODO: Share connection among model
	// TODO: Error handling
	conn, _ := dbr.Open(cxt.GetConf().DBType, cxt.GetConf().GetDSN(), nil)
	sess := conn.NewSession(nil)
	defer sess.Close()
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
