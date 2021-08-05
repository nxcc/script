package main

import (
	"fmt"
	"strings"

	"github.com/nxcc/script"
)

func main() {
	soc, _ := script.File("fruit.txt").SliceOfColumns()
	for _, line := range soc {
		fmt.Print(strings.Join(line, " | "), "\n")
	}
}
