package model

import (
	"database/sql"
	"log"
	"testing"

	"github.com/gchaincl/dotsql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/suite"
	"github.com/westlab/door-api/conf"
	"github.com/westlab/door-api/context"
)

type WordTestSuite struct {
	suite.Suite
}

func TestWordTestSuite(t *testing.T) {
	suite.Run(t, new(WordTestSuite))
}

func (s *WordTestSuite) SetupTest() {
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

	_, err = dsql.Exec(db, "drop-word-table")
	if err != nil {
		log.Fatal(err)
	}

	_, err = dsql.Exec(db, "create-word-table")
	if err != nil {
		log.Fatal(err)
	}

	_, err = dsql.Exec(db, "insert-word-rows")
	if err != nil {
		log.Fatal(err)
	}
}

func (s *WordTestSuite) TearDonwTest() {
}

func (s *WordTestSuite) TestWordSave() {
	var name string
	var count int64
	var w *Word

	name = "INTEROP2016TWS"
	count = 3
	w = NewWord(name, count)
	w.Save()
	count = 5
	w = NewWord(name, count)
	w.Save()

	counts := GetWordCount(int64(0))
	var savedFlag bool
	for _, c := range counts {
		if c.Name == "INTEROP2016TWS" {
			savedFlag = true
			s.Equal(8, int(c.Count))
		}
	}
	s.Equal(true, savedFlag)
}
