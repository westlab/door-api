package model

import (
	"strconv"

	// lib for mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/westlab/door-api/conf"
)

// Meta is meta information for door api
type Meta struct {
	Name      string       `db:"name" json:"name"`
	Value     string       `db:"value" json:"value"`
	CreatedAt dbr.NullTime `db:"created_at" json:"created_at"`
}

// ToInt converts value to int64
func (m *Meta) ToInt() int64 {
	i, _ := strconv.Atoi(m.Value)
	return int64(i)
}

// ToBool converts value to bool
func (m *Meta) ToBool() bool {
	b, _ := strconv.ParseBool(m.Value)
	return b
}

// NewMeta creates a new Meta
func NewMeta(name string, value string) Meta {
	// TODO: implement add Meta using dbr
	return Meta{}
}

// Select from table
func SelectSingleMeta(name string) *Meta {
	var m *Meta

	conf := conf.GetConf()
	conn, _ := dbr.Open(conf.DBType, conf.GetDSN(), nil)
	sess := conn.NewSession(nil)
	sess.Select("name", "value", "created_at").From("meta").Where("name = ?", name).LoadStruct(&m)
	return m
}
