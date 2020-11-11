package main

import (
	"flag"
	"time"

	"github.com/polevpn/elog"
)

func main() {
	flag.Parse()
	defer elog.Flush()
	elog.SetMode(elog.WITH_NO_FILE_LINE)
	for i := 0; i < 100; i++ {
		go func() {
			elog.Info("hello", "world")
		}()
	}
	elog.Fatal("xxxxxxxxxxxxxxxxxxxxxxxxxx")
	time.Sleep(time.Second)
}
