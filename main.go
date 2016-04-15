package main

import (
	"github.com/labstack/echo/engine/standard"
	"github.com/westlab/door-api/route"
)

func main() {
	router := route.Init()
	router.Run(standard.New(":8080"))
}
