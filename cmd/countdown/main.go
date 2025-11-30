package main

import (
	"os"
	"time"

	"github.com/dksifoua/golearn/pkg/mocking"
)

func main() {
	if err := mocking.Countdown(os.Stdout, time.Sleep, 20); err != nil {
		panic("Something went wrong!")
	}
}
