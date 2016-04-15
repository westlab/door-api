package main

import (
	"github.com/westlab/door-api/route"
)

func main() {
	router := route.Init()
	router.Run(":8080")
}
