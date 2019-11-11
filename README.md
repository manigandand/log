# log

[![Go Report](https://goreportcard.com/badge/github.com/manigandand/log)](https://goreportcard.com/report/github.com/manigandand/log)
[![GolangCI](https://golangci.com/badges/github.com/manigandand/log.svg)](https://golangci.com/r/github.com/manigandand/log)
[![License](https://img.shields.io/badge/license-MIT%20License-blue.svg)](https://github.com/manigandand/log/blob/master/LICENSE)
[![Build Status](https://img.shields.io/travis/manigandand/log?logo=travis&style=for-the-badge)](https://travis-ci.org/manigandand/log)
[![Coverage Status](https://img.shields.io/codecov/c/gh/manigandand/log.svg?logo=codecov&style=for-the-badge)](https://codecov.io/gh/manigandand/log)
[![](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://godoc.org/github.com/manigandand/log)

Yet another Go Logger

log is simple extended version of standard 'log' package based on logLevel
most of the concepts inspired from [glog](https://godoc.org/github.com/golang/glog) and [tracelog](https://github.com/goinggo/tracelog) these packages are huge and complex.

Hence we writing our own log package with as simple as possible and extensible to third party logging systems such as Papertrail, Loggly, LogDNA.

```golang
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

```

```Shell
INFO: 2019/10/07 18:37:20 main.go:35: log info level example
INFO: 2019/10/07 18:37:20 main.go:36: time---1570453640
INFO: 2019/10/07 18:37:20 main.go:37: log info level example
```

```Shell
WARNING: 2019/10/07 18:37:20 main.go:39: log info level example
WARNING: 2019/10/07 18:37:20 main.go:40: time---1570453640
WARNING: 2019/10/07 18:37:20 main.go:41: log info level example
```

```Shell
ERROR: 2019/10/07 18:37:20 main.go:43: log info level example
ERROR: 2019/10/07 18:37:20 main.go:44: time---1570453640
ERROR: 2019/10/07 18:37:20 main.go:45: log info level example
```

```Shell
FATAL: 2019/10/07 18:37:20 main.go:47: log info level example
exit status 1
```

```Shell
PANIC: 2019/10/07 18:37:20 main.go:47: log info level example
exit status 1
```

> TODO:

-   Context based logging.
-   github.com/sirupsen/logrus
-   github.com/uber-go/zap

## Licence

MIT
