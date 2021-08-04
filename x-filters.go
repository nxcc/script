package script

import (
	"bytes"
	"io/ioutil"
	"strings"
	"text/template"
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
