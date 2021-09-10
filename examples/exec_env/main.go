package main

import (
	"os"

	"github.com/nxcc/script"
)

func main() {

	os.Setenv("FOO", "bar42")

	script.
		SetEnv(script.Env{"HELLO": "Bartman"}).
		Exec(
			`bash -c 'echo "hello ${HELLO:-stranger} (FOO=$FOO)"'`,
		).Stdout()
}
