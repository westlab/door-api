package model

type Count struct {
	Name  string `db:"name" json:"name"`
	Count int64  `db:"count" json:"count"`
}
