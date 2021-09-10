package script

import "fmt"

func Execf(s string, a ...interface{}) *Pipe {
	return NewPipe().Exec(fmt.Sprintf(s, a...))
}

func SetEnv(env Env) *Pipe {
	return NewPipe().SetEnv(env)
}
