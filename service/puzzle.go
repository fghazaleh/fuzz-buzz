package service

import (
	"strconv"
	"strings"
)

type FuzzBuzzer interface {
	IsValid(int) bool
	GetOutput() string
}

type puzzle struct {
	collection []FuzzBuzzer
}

func NewPuzzle() *puzzle {
	return &puzzle{[]FuzzBuzzer{}}
}

func (p *puzzle) Add(s FuzzBuzzer) {
	p.collection = append(p.collection, s)
}

func (p *puzzle) SetupDefault() {
	p.Add(NewBuzz())
	p.Add(NewFuzz())
}

func (p *puzzle) Handle(m int) string {
	output := ""
	var out string
	for i := 1; i <= m; i++ {
		out = ""
		for _, service := range p.collection {
			if service.IsValid(i) {
				out += service.GetOutput()
			}
		}

		if out != "" {
			output += out + ","
		} else {
			output += strconv.Itoa(i) + ","
		}
	}
	return strings.TrimRight(output, ",")
}
