package model

import (
	// lib for mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
	"strconv"
)

// Meta is meta information for door api
type Meta struct {
	ID    int64  `db:"id" json:"json"`
	Name  string `db:"name" json:"name"`
	Value string `db:"value" json:"value"`
}

// ToInt converts value to int64
func (m *Meata) ToInt() int64 {
	i, _ := strconv.Atoi(m.Value)
	return i
}

// ToBool converts value to bool
func (m *Meata) ToBool() bool {
	b, _ := strconv.ParseBool(m.Value)
	return b
}

// NewMeta creates a new Meata
func NewMeta(name string, value string) Meta {
	// TODO: implement add Meta using dbr
}
