package main

import (
	"fmt"

	"github.com/richardimaoka/goxample/greet"
)

func main() {
	greet.Hello()
	h := greet.NewHidden()
	fmt.Printf("%v\n", h.MethodOnHidden())

	fmt.Println(greet.Shark)

	oct := greet.Octopus{
		Name:  "Jesse",
		Color: "orange",
	}

	fmt.Println(oct.String())
}
