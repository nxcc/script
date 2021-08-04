package main

import (
	"github.com/nxcc/script"
)

func main() {
	script.Exec("bash -c 'echo hello world'").Stdout()
}
