package script

import "fmt"

func Execf(s string, a ...interface{}) *Pipe {
	return NewPipe().Exec(fmt.Sprintf(s, a...))
}
