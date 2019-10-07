package main

import (
	"flag"
	"io"
	"log"

	"github.com/manigandand/log"
)

var (
	logLevel = flag.String("log", "INFO",
		"log-level for the app valid choice INFO, WARNING, ERROR, FATAL, PANIC")
	papertrailAddr = flag.String("papertrail", "",
		"with valid papertrail address starts logging on papertrail.")
)

func main() {
	flag.Parse()

	// init log
	var multiWriter io.Writer
	if *papertrailAddr != "" {
		multiWriter = &log.Papertrail{Address: *papertrailAddr}
	}
	l, ok := log.Levels[*logLevel]
	if !ok {
		log.Fatal("logx: invalid log-level")
	}
	log.Init(l, multiWriter)
}
