package model

import (
	"database/sql"
	"fmt"
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
func (m *Meta) ToInt() (int64, error) {
	i, err := strconv.Atoi(m.Value)
	return int64(i), err
}

// ToBool converts value to bool
func (m *Meta) ToBool() (bool, error) {
	b, err := strconv.ParseBool(m.Value)
	return b, err
}

func CreateOrUpdateMeta(key string, value string) (sql.Result, error) {
	var m *Meta
	conf := conf.GetConf()
	conn, _ := dbr.Open(conf.DBType, conf.GetDSN(), nil)
	sess := conn.NewSession(nil)
	sess.Select("name", "value", "created_at").From("meta").Where("name = ?", name).Load(&m)
	if m != nil {
		return sess.InsertInto("meta").Columns("name", "value").Values(key, value).Exec()
	}
	return sess.Update("meta").Set("name", key).Set("value", value).Exec()
}

// MarshalJSON overrides MarshalJSON
func (m *Meta) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(
		"{\"name\":\"%s\", \"value\":\"%s\", \"created_at\":\"%s\"}",
		m.Name,
		m.Value,
		m.CreatedAt.Time.Format("2006-01-02 15:04:05"))), nil
}

// NewMeta creates a new Meta
func NewMeta(name string, value string) Meta {
	// TODO: implement add Meta using dbr
	return Meta{}
}

// SelectSingleMeta selects meta by name
func SelectSingleMeta(name string) *Meta {
	var m *Meta

	conf := conf.GetConf()
	conn, _ := dbr.Open(conf.DBType, conf.GetDSN(), nil)
	sess := conn.NewSession(nil)
	sess.Select("name", "value", "created_at").From("meta").Where("name = ?", name).Load(&m)
	return m
}
