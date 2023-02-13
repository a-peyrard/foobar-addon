package main

import (
	"fmt"
	"github.com/a-peyrard/addon-manager/process"
)

type Process struct{}

func (p *Process) Run() (err error) {
	fmt.Printf("\n== foooooooobaaaaaaaar v2.0.0! ==\n\n")

	return
}

func NewProcess() process.Process {
	return &Process{}
}
