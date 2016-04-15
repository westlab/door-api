package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
	"strconv"
)

type Meta struct {
	Id    int64  `db:"id" json:"json"`
	Name  string `db:"name" json:"name"`
	Value string `db:"value" json:"value"`
}

func (m *Meata) ToInt() int64 {
	i, _ := strconv.Atoi(m.Value)
	return i
}

func (m *Meata) ToBool() bool {
	b, _ := strconv.ParseBool(m.Value)
	return b
}

func NewMeta(name string, value string) Meta {
	// TODO: implement add Meta using dbr
}
