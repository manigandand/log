package main

import (
	"flag"
	"io"
	"time"

	syslog "log"

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
		syslog.Fatal("log: invalid log-level")
	}
	log.Init(l, multiWriter)

	// ex
	log.Info("log info level example")
	log.Infof("%s---%d", "time", time.Now().Unix())
	log.Infoln("log info level example")

	log.Warning("log info level example")
	log.Warningf("%s---%d", "time", time.Now().Unix())
	log.Warningln("log info level example")

	log.Error("log info level example")
	log.Errorf("%s---%d", "time", time.Now().Unix())
	log.Errorln("log info level example")

	log.Fatal("log info level example")
	log.Fatalf("%s---%d", "time", time.Now().Unix())
	log.Fatalln("log info level example")

	log.Panic("log info level example")
	log.Panicf("%s---%d", "time", time.Now().Unix())
	log.Panicln("log info level example")
}
