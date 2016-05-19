package conf

import (
	"fmt"
	"log"

	"github.com/pelletier/go-toml"
)

var conf *Config

// Config for door application
type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int64
	DBType     string
	DBName     string
	AppPort    int64
	AppDebug   bool
	Sockets    []string
}

// New generate Config singleton
func New(tomlFile string) *Config {
	tomlConf, err := toml.LoadFile(tomlFile)
	// Return error
	if err != nil {
		log.Println("Parsing config was failed ", err.Error())
	}

	// Convert interface to slice
	// handle nil case
	confSocket := tomlConf.Get("app.sockets").([]interface{})
	sockets := make([]string, len(confSocket), len(confSocket))
	for i, e := range confSocket {
		sockets[i] = e.(string)
	}

	conf = &Config{
		DBUser:     tomlConf.Get("db.user").(string),
		DBPassword: tomlConf.Get("db.password").(string),
		DBHost:     tomlConf.Get("db.host").(string),
		DBPort:     tomlConf.Get("db.port").(int64),
		DBType:     tomlConf.Get("db.type").(string),
		DBName:     tomlConf.Get("db.dbname").(string),
		AppPort:    tomlConf.Get("app.port").(int64),
		AppDebug:   tomlConf.Get("app.debug").(bool),
		Sockets:    sockets,
	}
	return conf
}

// GetConf returns Config singleton
func GetConf() *Config {
	return conf
}

// GetDSN returns DSN for db connection
func (c *Config) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName,
	)
}
