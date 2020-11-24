package elog

import "testing"

var tlog *EasyLogger

func TestGetLogger(t *testing.T) {
	tlog = GetLogger()
	defer tlog.Flush()

	tlog.Info("Hello")
	tlog.Info("Hello")
}

func TestElog(t *testing.T) {
	defer logger.Flush()

	logger.Info("Hello")
	logger.Info("Hello")
}
