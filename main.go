package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/labstack/echo/engine/standard"

	"github.com/westlab/door-api/conf"
	"github.com/westlab/door-api/context"
	"github.com/westlab/door-api/manager"
	"github.com/westlab/door-api/route"
)

func main() {
	confFile := flag.String("f", "", "path to the config file")
	flag.Parse()
	log.Println(*confFile)

	conf := conf.New(*confFile)
	cxt := context.NewContext(conf)

	// Start Job manager
	manager := manager.NewHTTPJobManager(cxt)
	manager.Start()

	// Start Server
	router := route.Init(cxt)
	ipPort := fmt.Sprintf(":%d", cxt.GetConf().AppPort)
	router.Run(standard.New(ipPort))
}
