package main

import (
	"github.com/lib4dev/cli"
	_ "github.com/micro-plat/hycli/mgr"
	_ "github.com/micro-plat/hycli/mob"
	"github.com/micro-plat/lib4go/logger"
)

func main() {
	logger.Pause()
	var app = cli.New(cli.WithVersion("0.1.0"))
	app.Start()
}
