package main

import (
	"os"

	"github.com/nxcc/script"
)

func main() {
	script.Stdin().Match(os.Args[1]).Stdout()
}
