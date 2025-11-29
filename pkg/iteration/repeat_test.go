package iteration

import "testing"

func assertEqual(t testing.TB, expected, actual string) {
	t.Helper()

	if actual != expected {
		t.Errorf("Expected: %s - Actual %s", expected, actual)
	}
}

func TestRepeat(t *testing.T) {
	expected := "aaaaa"
	actual := Repeat("a", 5)

	assertEqual(t, expected, actual)
}

func TestRepeatOptimal(t *testing.T) {
	expected := "aaaaa"
	actual := RepeatOptimal('a', 5)

	assertEqual(t, expected, actual)
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 20)
	}
}

func BenchmarkRepeatOptimal(b *testing.B) {
	for b.Loop() {
		RepeatOptimal('a', 20)
	}
}
