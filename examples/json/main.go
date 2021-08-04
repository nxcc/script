package main

import (
	"fmt"

	"github.com/nxcc/script"
)

type Foo struct {
	Bar int `json:"bar"`
}

func main() {
	foo := Foo{}
	err := script.Echo(`{"bar":42}`).JSON(&foo)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", foo)
}
