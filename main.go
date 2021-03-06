package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/engine/standard"

	"github.com/westlab/door-api/conf"
	"github.com/westlab/door-api/context"
	"github.com/westlab/door-api/manager"
	"github.com/westlab/door-api/route"
)

func main() {
	confFile := flag.String("f", "", "path to the config file")
	flag.Parse()

	conf := conf.New(*confFile)
	cxt := context.NewContext(conf)

	// Start Job manager
	if cxt.GetConf().HTTPReconstructor {
		httpManager := manager.NewHTTPJobManager(cxt)
		httpManager.Start()
	}

	// Start Browsing time manager
	if cxt.GetConf().BrowsingTimer {
		browsingManager := manager.NewBrowsingTimeManager(cxt)
		browsingManager.Start()
	}

	// Start HTML analyzer manager
	if cxt.GetConf().HTMLAnalyzer {
		htmlAnalyzeManager := manager.NewHTMLAnalyzerManager(cxt)
		htmlAnalyzeManager.Start()
	}

	// Start Server
	router := route.Init(cxt)
	ipPort := fmt.Sprintf(":%d", cxt.GetConf().AppPort)
	go router.Run(standard.New(ipPort))

	// Clean up sockets
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-c
	log.Println("Clean up")
	for _, sock := range cxt.GetConf().Sockets {
		os.Remove(sock)
	}
}
