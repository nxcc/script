package main

import (
	"github.com/nxcc/script"
)

func init() { script.DebugExec = true }

func main() {
	// this is just a silly example, the same could easier be accomplished with ".Match"
	script.Echo("one\ntwo\nthree\nfour\n").Execf("grep -v %v", "three").Stdout()
}
