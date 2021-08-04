package main

import (
	"github.com/nxcc/script"
)

// This program prints out the names of all files in the current directory.

func main() {
	script.ListFiles("*").ExecForEach("echo {{.}}").Stdout()
}
