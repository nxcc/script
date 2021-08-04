package main

import "github.com/nxcc/script"

func main() {

	script.
		Echo("one\n{{.x}}\nthree\n").
		Template(map[string]string{"x": "two"}).
		Stdout()

	script.
		File("template.tmpl").
		Template(map[string]string{"value": "two"}).
		Stdout()
}
