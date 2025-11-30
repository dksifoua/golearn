package mocking

import (
	"bytes"
	"testing"
	"time"
)

var sleepMock SleepFn = func(d time.Duration) {
	var count int64
	for count = 0; count < d.Milliseconds(); count++ {
	}
}

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}

	expected := `5
4
3
2
1
Go!
`
	if err := Countdown(&buffer, sleepMock, 5); err != nil {
		t.Errorf("Expected: %s - Actual: nil", expected)
	}

	actual := buffer.String()
	if actual != expected {
		t.Errorf("Expected: %s - Actual: %s", expected, actual)
	}
}
