package logger

import "testing"

/*
File name    : logger_test.go
Author       : miaoyc
Create Date  : 2023/7/27 15:11
Update Date  : 2023/7/27 17:11
Description  :
*/

func init() {
	// log level: panic, fatal, error, warn/warning, info, debug, trace
	Setup("./run.log", "debug", 24, 24)
}

func TestInfo(t *testing.T) {
	Trace("trace world!")
	Tracef("trace world! %s", "miaoyc")
	Debug("debug world!")
	Debugf("debug world! %s", "miaoyc")
	Info("info world!")
	Infof("info world! %s", "miaoyc")
	Warn("warn world!")
	Warnf("warn world! %s", "miaoyc")
	Error("error world!")
	Errorf("error world! %s", "miaoyc")
	Fatal("fatal world!")
	Fatalf("fatal world! %s", "miaoyc")
	Panic("panic world!")
	Panicf("panic world! %s", "miaoyc")
}
