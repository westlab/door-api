package model

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/gchaincl/dotsql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/westlab/door-api/conf"
	"github.com/westlab/door-api/context"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	createTestTable()
}

func teardown() {
	dropTestTable()
}

func createTestTable() {
	// load conf file
	conf := conf.New("../config.toml")
	context.NewContext(conf)
	cxt := context.GetContext()

	db, err := sql.Open(cxt.GetConf().DBType, cxt.GetConf().GetDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dsql, err := dotsql.LoadFromFile("../test/test_data.sql")
	if err != nil {
		log.Fatal(err)
	}

	// Drop Tables
	_, err = dsql.Exec(db, "drop-browsing-table")
	if err != nil {
		log.Fatal(err)
	}
	_, err = dsql.Exec(db, "drop-meta-table")
	if err != nil {
		log.Fatal(err)
	}

	// Create Tables
	_, err = dsql.Exec(db, "create-browsing-table")
	if err != nil {
		log.Fatal(err)
	}
	_, err = dsql.Exec(db, "create-meta-table")
	if err != nil {
		log.Fatal(err)
	}

	// Insert Rows
	_, err = dsql.Exec(db, "insert-browsing-rows")
	if err != nil {
		log.Fatal(err)
	}
	_, err = dsql.Exec(db, "insert-meta-rows")
	if err != nil {
		log.Fatal(err)
	}
}

func dropTestTable() {
	conf := conf.New("../config.toml")
	context.NewContext(conf)

	db, err := sql.Open(conf.DBType, conf.GetDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	dsql, err := dotsql.LoadFromFile("../test/test_data.sql")
	if err != nil {
		log.Fatal(err)
	}

	// Drop Tables
	_, err = dsql.Exec(db, "drop-browsing-table")
	if err != nil {
		log.Fatal(err)
	}
	_, err = dsql.Exec(db, "drop-meta-table")
	if err != nil {
		log.Fatal(err)
	}
}