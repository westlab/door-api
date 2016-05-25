package conf

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tomlFile *os.File

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

// createTempConfig is a helper function to generate toml format config file
func createTempConfig() {
	tomlFile, _ = ioutil.TempFile(os.TempDir(), "door-api")
	tomlData := `
		[db]
		user = 'interop'
		password = 'interop'
		host = 'localhost'
		port = 3306
		type = 'mysql'
		dbname = 'interop2016'

		[app]
		port = 8080
		debug = true
		sockets = ["/tmp/foo", "/tmp/bar", "/tmp/baz"]
        blackList = "/tmp/blacklist"
        words = "/tmp/words"
	`
	ioutil.WriteFile(tomlFile.Name(), []byte(tomlData), os.ModePerm)
}

func setup() {
	createTempConfig()
}

func teardown() {
	os.Remove(tomlFile.Name())
}

func TestToConf(t *testing.T) {
	conf := New(tomlFile.Name())
	assert.Equal(t, conf.DBType, "mysql")
	assert.Equal(t, conf.DBUser, "interop")
	assert.Equal(t, conf.DBPassword, "interop")
	assert.Equal(t, conf.DBPort, int64(3306))
	assert.Equal(t, conf.AppPort, int64(8080))
	assert.Equal(t, conf.AppDebug, true)
	assert.Equal(t, conf.Sockets, []string{"/tmp/foo", "/tmp/bar", "/tmp/baz"})
	assert.Equal(t, conf.BlackList, "/tmp/blacklist")
	assert.Equal(t, conf.WordsPath, "/tmp/words")
}

func TestGetDSN(t *testing.T) {
	conf := New(tomlFile.Name())
	assert.Equal(t, conf.GetDSN(), "interop:interop@tcp(localhost:3306)/interop2016")
}
