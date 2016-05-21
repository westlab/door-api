package job

import (
	"net"
	"os"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/westlab/door-api/common"
)

var socketFile string

func createTempSocket() {
	socketFileDir := os.TempDir()
	socketFile = path.Join(socketFileDir, common.RandomStr())
}

func setup() {
	createTempSocket()
}

func teardown() {
	os.Remove(socketFile)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func createClient(socket string) net.Conn {
	c, err := net.Dial("unix", socket)
	if err != nil {
		panic(err)
	}
	return c
}

func TestBasic(t *testing.T) {
	ch := make(chan *string)
	reciever := NewDoorReciever(socketFile, ch)
	go reciever.Start()

	// Wait to make sure socket is up
	time.Sleep(100 * time.Millisecond)
	client := createClient(socketFile)

	client.Write([]byte("1"))
	d := <-ch
	assert.Equal(t, "1", *d)
}
