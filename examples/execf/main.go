package main

import (
	"github.com/nxcc/script"
)

func main() {
	script.Execf("bash -c 'echo hello %v'", "world").Stdout()
}
