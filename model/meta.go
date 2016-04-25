package model

import (
	"strconv"
	"time"
)

// Meta is meta information for door api
type Meta struct {
	Name      string    `db:"name" json:"name"`
	Value     string    `db:"value" json:"value"`
	CreatedAt time.Time `db:created_at json:"created_at"`
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

// NewMeta creates a new Meata
func NewMeta(name string, value string) Meta {
	// TODO: implement add Meta using dbr
	return Meta{}
}
