package elog_test

import (
	"fmt"
	"testing"

	"github.com/polevpn/elog"
)

var tlog *elog.EasyLogger

func TestGetLogger(t *testing.T) {
	tlog = elog.GetLogger()
	defer tlog.Flush()

	tlog.Info("Hello")
	tlog.Info("Hello")
}

func TestElog(t *testing.T) {
	defer elog.Flush()

	elog.Info("Hello", "xxxxx")
	elog.Info("Hello", "xxxxx")
	elog.Infof("%d,%v,%s", 1, "xx", "xxxxxxx")
}

func TestSprint(t *testing.T) {
	data := fmt.Sprint("xxxx", "xxxxx", "xxxxx")
	fmt.Print(data, "sdsdsd")
}
