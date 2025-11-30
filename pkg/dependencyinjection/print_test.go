package dependencyinjection

import (
	"bytes"
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	name := "Dimitri"

	err := Greet(&buffer, name)
	if err != nil {
		t.Errorf("Expected: nil - Actual %s", err)
	}

	expected := fmt.Sprintf("Hello, %s!\n", name)
	actual, err := buffer.ReadString('\n')
	if err != nil {
		t.Errorf("Expected: nil - Actual %s", err)
	}
	if actual != expected {
		t.Errorf("Expected: %s - Actual %s", expected, actual)
	}
}
