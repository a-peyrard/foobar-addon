package main

import "fmt"

type Process struct{}

func (p *Process) Run() (err error) {
	fmt.Printf("\n== foooooooobaaaaaaaar v1.0.0! ==\n\n")

	return
}

func NewProcess() *Process {
	return &Process{}
}
