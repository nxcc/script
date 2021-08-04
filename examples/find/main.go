package main

import (
	"os"

	"github.com/nxcc/script"
)

func main() {
	listPath := "."
	if len(os.Args) > 1 {
		listPath = os.Args[1]
	}
	script.FindFiles(listPath).Stdout()
}
