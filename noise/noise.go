package main

import (
	"fmt"
	"log"
	"log/syslog"
	"time"
)

var (
	logger *syslog.Writer
	err    error
)

func init() {
	if logger, err = syslog.Dial("tcp", "syslog:514", syslog.LOG_INFO, "elk"); err != nil {
		log.Fatalf("%v", err)
	}
}

func main() {
	defer logger.Close()

	tiker := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-tiker.C:
			go func() {
				logger.Info(fmt.Sprintf("=> clock is ticking ... %v", time.Now()))
			}()
		}
	}
}
