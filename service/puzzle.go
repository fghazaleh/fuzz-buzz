package service

import (
	"strconv"
	"strings"
)

type FuzzBuzzer interface {
	IsValid(int) bool
	GetOutput() string
}

type cc []FuzzBuzzer

type puzzle struct {
	collection cc
}

func NewPuzzle() *puzzle {
	return &puzzle{cc{}}
}

func (p *puzzle) Add(s FuzzBuzzer) {
	p.collection = append(p.collection, s)
}

func (p puzzle) Handle(m int) string {
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
