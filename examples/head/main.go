package main

import (
	"log"
	"os"
	"strconv"

	"github.com/nxcc/script"
)

func main() {
	lines, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	script.Stdin().First(lines).Stdout()
}
