package bale

import (
	"flag"
	"testing"

	"github.com/cee-dub/bale/glog"
)

func init() {
	flag.Parse()
}

func TestGlogSatisfiesLogger(t *testing.T) {
	var log Logger
	log = glog.L
}
