package manager

import (
	"database/sql"
	"log"
	"net"
	"os"
	"testing"
	"time"

	"github.com/gchaincl/dotsql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"

	"github.com/stretchr/testify/assert"
	"github.com/westlab/door-api/conf"
	"github.com/westlab/door-api/context"
)

const (
	socket = "/tmp/door1"
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	createTable()
}

func teardown() {

}

func createClient(socket string) net.Conn {
	c, err := net.Dial("unix", socket)
	if err != nil {
		panic(err)
	}
	return c
}

func createTable() {
	conf := conf.New("../config.toml")
	os.Remove(socket)
	conf.Sockets = []string{socket}
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

	_, err = dsql.Exec(db, "drop-browsing-table")
	if err != nil {
		log.Fatal(err)
	}

	_, err = dsql.Exec(db, "create-browsing-table")
	if err != nil {
		log.Fatal(err)
	}
}

func TestHTTPReconstructionFromDoor(t *testing.T) {
	cxt := context.GetContext()

	manager := NewHTTPJobManager(cxt)
	manager.Start()

	// Wait to make sure socket is up
	time.Sleep(100 * time.Millisecond)
	client := createClient(socket)

	client.Write([]byte("1.1.1.1,2.2.2.2,12345,80,2016-04-01 00:00:00,GET, /home HTTP 1.1DEND"))
	client.Write([]byte("1.1.1.1,2.2.2.2,12345,80,2016-04-01 00:00:00,Host:, google.com\r\n From: user@example.comDEND"))
	client.Write([]byte("2.2.2.2,1.1.1.1,80,12345,2016-04-01 00:00:00,Content-Type:, text/htmlDEND"))
	client.Write([]byte("2.2.2.2,1.1.1.1,80,12345,2016-04-01 00:00:00,<title, >WestLab</titile><body></body>DEND"))

	conn, _ := dbr.Open(cxt.GetConf().DBType, cxt.GetConf().GetDSN(), nil)
	sess := conn.NewSession(nil)

	time.Sleep(100 * time.Millisecond)
	var r int64
	sess.Select("count(id)").From("browsing").Load(&r)

	assert.Equal(t, int64(1), r)
}
