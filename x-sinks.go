package script

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

func (p *Pipe) JSON(obj interface{}) error {
	if p == nil || p.Error() != nil {
		return p.Error()
	}
	res, err := ioutil.ReadAll(p.Reader)
	if err != nil {
		p.SetError(err)
		return err
	}
	err = json.Unmarshal(res, obj)
	if err != nil {
		p.SetError(err)
		return err
	}
	return nil
}

func (p *Pipe) End() (stdout, stderr string, err error) {
	if p == nil || p.Error() != nil {
		err := p.Error()
		p.SetError(nil)
		stderr, _ := p.String()
		return "", stderr, err
	}
	output, _ := p.String()
	return output, "", nil
}

func (p *Pipe) SliceOfColumns() ([][]string, error) {
	if p == nil || p.Error() != nil {
		return nil, p.Error()
	}
	result := [][]string{}
	p.EachLine(func(line string, out *strings.Builder) {
		result = append(result, strings.Fields(line))
	})
	return result, p.Error()
}
