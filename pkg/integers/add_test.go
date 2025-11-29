package integers

import "testing"

func TestAdd(t *testing.T) {
	expected := 5
	actual := Add(2, 3)

	if actual != expected {
		t.Errorf("Expected: %d - Actual: %d", expected, actual)
	}
}
