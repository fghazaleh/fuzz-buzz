package main

import (
	"fmt"
	"fuzz-buzz/service"
)

func main() {

	c := service.NewPuzzle()
	c.SetupDefault()
	//or
	//c.Add(service.NewBuzz())
	//c.Add(service.NewFuzz())

	fmt.Println(c.Handle(20))

	fmt.Println("Done !")
}
