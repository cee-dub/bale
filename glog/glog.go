package glog

import (
	"github.com/golang/glog"
)

type Log struct{}

var L = Log{}

func (l Log) Info(args ...interface{}) {
	glog.Info(args...)
}

func (l Log) Infof(format string, args ...interface{}) {
	glog.Infof(format, args...)
}

func (l Log) Warn(args ...interface{}) {
	glog.Warning(args...)
}

func (l Log) Warnf(format string, args ...interface{}) {
	glog.Warningf(format, args...)
}

func (l Log) Error(args ...interface{}) {
	glog.Error(args...)
}

func (l Log) Errorf(format string, args ...interface{}) {
	glog.Errorf(format, args...)
}

func (l Log) Fatal(args ...interface{}) {
	glog.Fatal(args...)
}

func (l Log) Fatalf(format string, args ...interface{}) {
	glog.Fatalf(format, args...)
}

func (l Log) flush() {
	glog.Flush()
}
