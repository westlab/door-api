package model

import (
	"strconv"
)

// Meta is meta information for door api
type Meta struct {
	ID    int64  `db:"id" json:"json"`
	Name  string `db:"name" json:"name"`
	Value string `db:"value" json:"value"`
}

// ToInt converts value to int64
func (m *Meta) ToInt() int {
	i, _ := strconv.Atoi(m.Value)
	return i
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
