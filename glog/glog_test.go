package glog

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Severities, helper functions & tests borrowed from glog. Original Copyright
// 2013 Google Inc.

// severity identifies the sort of log: info, warning etc.
type severity int

const (
	infoLog severity = iota
	warningLog
	errorLog
	fatalLog
	numSeverity = 4
)

var severityName = []string{
	infoLog:    "INFO",
	warningLog: "WARNING",
	errorLog:   "ERROR",
	fatalLog:   "FATAL",
}

// contents returns the specified log value as a string.
func contents(s severity) string {
	p := filepath.Join(os.TempDir(), "glog.test."+severityName[s])
	b, _ := ioutil.ReadFile(p)
	return string(b)
}

// contains reports whether the string is contained in the log.
func contains(s severity, str string) bool {
	return strings.Contains(contents(s), str)
}

// Test that Info works as advertised.
func TestInfo(t *testing.T) {
	L.Info("test")
	L.flush()
	if !contains(infoLog, "test") {
		t.Errorf("Info has wrong character: %q", contents(infoLog))
	}
	if !contains(infoLog, "test") {
		t.Error("Info failed")
	}
}

// Test that Warn works as advertised.
func TestWarn(t *testing.T) {
	L.Warn("test")
	L.flush()
	if !contains(infoLog, "test") {
		t.Errorf("Warn has wrong character: %q", contents(infoLog))
	}
	if !contains(infoLog, "test") {
		t.Error("Warn failed")
	}
}

// Test that Error works as advertised.
func TestError(t *testing.T) {
	L.Error("test")
	L.flush()
	if !contains(infoLog, "test") {
		t.Errorf("Error has wrong character: %q", contents(infoLog))
	}
	if !contains(infoLog, "test") {
		t.Error("Error failed")
	}
}
