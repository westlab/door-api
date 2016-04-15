package model

import (
	// lib for mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
)

// Word is a model of word to track frequency
type Word struct {
	ID    int64  `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Count int64  `db:"count" json:"count"`
}

// NewWord creates a new Word
func NewWord(name string, count int64) Word {
	return Word{}
}
