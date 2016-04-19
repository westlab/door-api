package conf

import (
	"fmt"
)

const (
	// USER is username for db
	USER string = "interop"
	// PASSWORD is password for db
	PASSWORD string = "interop"
	// DB is name of db
	DB string = "interop2016"
	// HOST is host of db
	HOST string = "localhost"
	// PORT is  port number of db
	PORT string = "3306"
	// DBTYPE allows mysql and postgres
	DBTYPE string = "mysql"
)

// Return DSN for local connection
func GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		USER, PASSWORD, HOST, PORT, DB,
	)
}

func GetDBType() string {
	return DBTYPE
}
