package greet

import "fmt"

var Shark = "Sammy"

type Octopus struct {
	Name  string
	Color string
}

type Greet struct {
	a string
	b string
}

type hiddenStruct struct {
	a int
}

func NewHidden() *hiddenStruct {
	return &hiddenStruct{10}
}

func (h *hiddenStruct) MethodOnHidden() int {
	return h.a + 10
}

func (o Octopus) String() string {
	return fmt.Sprintf("The octopus's name is %q and is the color %s.", o.Name, o.Color)
}

func Hello() {
	fmt.Println("Hello, World!")
}
