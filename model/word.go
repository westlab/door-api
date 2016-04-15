package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
)

type Word struct {
	Id    int64  `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Count int64  `db:"count" json:"count"`
}

func NewWord(name string, count int64) {

}
