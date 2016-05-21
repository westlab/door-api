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

type MetaTestSuite struct {
	suite.Suite
}

func TestMetaTestSuite(t *testing.T) {
	suite.Run(t, new(MetaTestSuite))
}

func (s *MetaTestSuite) SetupTest() {
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

	_, err = dsql.Exec(db, "drop-meta-table")
	if err != nil {
		log.Fatal(err)
	}

	_, err = dsql.Exec(db, "create-meta-table")
	if err != nil {
		log.Fatal(err)
	}

	_, err = dsql.Exec(db, "insert-meta-rows")
	if err != nil {
		log.Fatal(err)
	}
}

func (s *MetaTestSuite) TearDonwTest() {
}

func (s *MetaTestSuite) TestSelectSingleMeta() {
	m := SelectSingleMeta("door-version")
	s.Equal("door-version", m.Name)
	s.Equal("1.0", m.Value)
	s.Equal("2016-04-01 00:00:00", m.CreatedAt.Time.Format("2006-01-02 15:04:05"))
}
