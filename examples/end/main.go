package main

import (
	"fmt"

	"github.com/nxcc/script"
)

func main() {

	stdout, stderr, err := script.Exec("date").End()
	fmt.Printf("stdout: %q  stderr: %q  err: %v\n", stdout, stderr, err)

	stdout, stderr, err = script.Exec("rm file404.txt").End()
	fmt.Printf("stdout: %q  stderr: %q  err: %v\n", stdout, stderr, err)
}
