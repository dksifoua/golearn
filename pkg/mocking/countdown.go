package mocking

import (
	"fmt"
	"io"
	"iter"
	"time"
)

type SleepFn func(d time.Duration)

func Countdown(w io.StringWriter, sleep SleepFn, n int) (err error) {
	var line string
	for i := range countdownFrom(n) {
		if i == 0 {
			line = "Go!\n"
		} else {
			line = fmt.Sprintf("%d\n", i)
		}
		if _, err = w.WriteString(line); err != nil {
			break
		}
		sleep(1 * time.Second)
	}
	return
}

func countdownFrom(n int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := n; i >= 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}
