package main

import (
	"github.com/nxcc/script"
)

// This version of cat takes a list of files on the command line, and prints
// their concatenated contents.

func main() {
	script.Args().Concat().Stdout()
}
