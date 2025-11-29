package hello

import "testing"

func assertCorrectMessage(t testing.TB, expected, actual string) {
	t.Helper()

	if actual != expected {
		t.Errorf("Expected: %s - Actual %s", expected, actual)
	}
}

func TestHello(t *testing.T) {
	t.Run("Say hello with empty parameters", func(t *testing.T) {
		expected := "Hello, World!"
		actual := Hello("", "")

		assertCorrectMessage(t, expected, actual)
	})

	t.Run("Say hello to a specific person in english", func(t *testing.T) {
		expected := "Hello, Dimitri!"
		actual := Hello("English", "Dimitri")

		assertCorrectMessage(t, expected, actual)
	})

	t.Run("Say hello to a specific person in french", func(t *testing.T) {
		expected := "Bonjour, Dimitri!"
		actual := Hello("French", "Dimitri")

		assertCorrectMessage(t, expected, actual)
	})

	t.Run("Say hello to a specific person in spanish", func(t *testing.T) {
		expected := "Hola, Dimitri!"
		actual := Hello("Spanish", "Dimitri")

		assertCorrectMessage(t, expected, actual)
	})

	t.Run("Say hello to everyone in english", func(t *testing.T) {
		expected := "Hello, World!"
		actual := Hello("", "")

		assertCorrectMessage(t, expected, actual)
	})
}
