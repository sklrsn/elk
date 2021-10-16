package main

import (
	"fmt"
	"log"
	"log/syslog"
	"math/rand"
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
				max := rand.Intn(10)
				for i := 0; i < max; i++ {
					logger.Info(fmt.Sprintf("=> clock is ticking ... %v", time.Now()))
				}
			}()
		}
	}
}
