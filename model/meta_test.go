package model

import (
	"database/sql"
	"testing"

	"github.com/gchaincl/dotsql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"github.com/westlab/door-api/conf"
)

func newTestDB(t testing.TB) (*sql.DB, *dotsql.DotSql) {
	conf := conf.New("../config.toml")
	db, err := sql.Open(conf.DBType, conf.GetDSN())
	if err != nil {
		t.Fatalf("Open: %v", err)
	}

	dsql, err := dotsql.LoadFromFile("../test/test_data.sql")
	if err != nil {
		t.Fatalf("%v", err)
	}

	_, err = dsql.Exec(db, "drop-meta-table")
	if err != nil {
		t.Fatalf("%v", err)
	}

	_, err = dsql.Exec(db, "create-meta-table")
	if err != nil {
		t.Fatalf("%v", err)
	}

	return db, dsql
}

func closeDB(t testing.TB, db *sql.DB, dsql *dotsql.DotSql) {
	_, err := dsql.Exec(db, "drop-meta-table")
	if err != nil {
		t.Fatalf("%v", err)
	}

	err = db.Close()
	if err != nil {
		t.Fatalf("error closing DB: %v", err)
	}
}

func TestSelectSingleMeta(t *testing.T) {
	db, dsql := newTestDB(t)
	defer closeDB(t, db, dsql)

	// insert
	_, err := dsql.Exec(db, "insert-meta-rows")
	if err != nil {
		t.Fatalf("%v", err)
	}

	m := SelectSingleMeta("door-version")
	assert.Equal(t, "door-version", m.Name)
	assert.Equal(t, "1.0", m.Value)
	assert.Equal(t, "2016-04-01 00:00:00", m.CreatedAt.Time.Format("2006-01-02 15:04:05"))
}
