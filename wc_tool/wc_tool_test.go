package wc_tool

import (
	"flag"
	"testing"
)

func TestWc(t *testing.T) {

	arg := "-c"
	flag.StringVar(&arg, "", "", "")
	Wc()

}
