package script

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"text/template"

	"bitbucket.org/creachadair/shell"
)

func (p *Pipe) Template(data interface{}) *Pipe {
	if p == nil || p.Error() != nil {
		return p
	}
	res, err := ioutil.ReadAll(p.Reader)
	if err != nil {
		p.SetError(err)
		return p.WithError(err)
	}
	tpl, err := template.New("").Parse(string(res))
	if err != nil {
		return p.WithError(err)
	}
	output := strings.Builder{}
	err = tpl.Execute(&output, data)
	if err != nil {
		return p.WithError(err)
	}
	q := NewPipe()
	return q.WithReader(bytes.NewReader([]byte(output.String())))
}

func (p *Pipe) Execf(s string, a ...interface{}) *Pipe {
	if p == nil || p.Error() != nil {
		return p
	}
	q := NewPipe()
	cmdLine := fmt.Sprintf(s, a...)
	args, ok := shell.Split(cmdLine) // strings.Fields doesn't handle quotes
	if !ok {
		return p.WithError(fmt.Errorf("unbalanced quotes or backslashes in [%s]", cmdLine))
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = p.Reader
	output, err := cmd.CombinedOutput()
	if err != nil {
		q.SetError(err)
	}
	return q.WithReader(bytes.NewReader(output))
}
