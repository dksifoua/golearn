package main

import (
	"os"

	di "github.com/dksifoua/golearn/pkg/dependencyinjection"
)

func main() {
	if err := di.Greet(os.Stdout, "Dimitri"); err != nil {
		panic("Something went wrong!")
	}
}
