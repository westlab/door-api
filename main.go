package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/labstack/echo/engine/standard"
	"github.com/westlab/door-api/conf"
	"github.com/westlab/door-api/route"
)

func main() {
	confFile := flag.String("f", "", "path to the config file")
	flag.Parse()
	log.Println(*confFile)

	conf := conf.New(*confFile)
	router := route.Init(*conf)
	ipPort := fmt.Sprintf(":%d", conf.AppPort)
	router.Run(standard.New(ipPort))
}
