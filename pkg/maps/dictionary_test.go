package maps

import (
	"strings"
	"testing"
)

func TestDictionary(t *testing.T) {

	t.Run("Search a non existing word from the dictionary", func(t *testing.T) {
		dictionary, word := Dictionary{}, "Test"

		expected := ErrDictionaryWordNotFound(word)
		_, actual := dictionary.Search(word)
		if actual == nil {
			t.Errorf("Expected: %s - Actual: nil", expected)
		}
		if actual.Error() != expected.Error() {
			t.Errorf("Expected: %s - Actual: %s", expected, actual)
		}
	})

	t.Run("Search an existing word from the dictionary", func(t *testing.T) {
		word, definition := "Test", "This is the definition of the word Test."
		dictionary := Dictionary{word: definition}

		expected := definition
		actual, _ := dictionary.Search(word)
		if actual != expected {
			t.Errorf("Expected: %s - Actual: %s", expected, actual)
		}
	})

	t.Run("Add a non existing word to the dictionary", func(t *testing.T) {
		dictionary, word, definition := Dictionary{}, "Test", "This is the definition of the word Test."

		err := dictionary.Add(word, definition)
		if err != nil {
			t.Errorf("Expected: nil - Actual: %s", err)
		}
	})

	t.Run("Add an existing word to the dictionary", func(t *testing.T) {
		dictionary, word, definition := Dictionary{}, "Test", "This is the definition of the word Test."

		err := dictionary.Add(word, definition)
		if err != nil {
			t.Errorf("Expected: nil - Actual: %s", err)
		}

		expected := ErrDictionaryWordAlreadyExists(word)
		actual := dictionary.Add(word, definition)
		if actual == nil {
			t.Errorf("Expected: %s - Actual: nil", expected)
		}
		if actual.Error() != expected.Error() {
			t.Errorf("Expected: %s - Actual: %s", expected, actual)
		}
	})

	t.Run("Update a non existing word from the dictionary", func(t *testing.T) {
		dictionary, word := Dictionary{}, "Nil"

		expected := ErrDictionaryWordNotFound(word)
		actual := dictionary.Update(word, "This null value in Golang.")
		if actual == nil {
			t.Errorf("Expected: %s - Actual: nil", expected)
		}
		if actual.Error() != expected.Error() {
			t.Errorf("Expected: %s - Actual: %s", expected, actual)
		}
	})

	t.Run("Update an existing word from the dictionary", func(t *testing.T) {
		word, definition := "Test", "This is the original definition"
		dictionary := Dictionary{word: definition}

		updatedDefinition := strings.Replace(definition, "original", "updated", 1)
		err := dictionary.Update(word, updatedDefinition)
		if err != nil {
			t.Errorf("Expected: nil - Actual: %s", err)
		}

		expected := updatedDefinition
		actual, err := dictionary.Search(word)
		if err != nil {
			t.Errorf("Expected: nil - Actual: %s", err)
		}
		if actual != expected {
			t.Errorf("Expected: %s - Actual: %s", expected, actual)
		}
	})

	t.Run("Delete a non existing word from the dictionary", func(t *testing.T) {
		dictionary, word := Dictionary{}, "Test"

		expected := ErrDictionaryWordNotFound(word)
		actual := dictionary.Delete(word)
		if actual == nil {
			t.Errorf("Expected: %s - Actual: nil", expected)
		}
	})

	t.Run("Delete an existing word from the dictionary", func(t *testing.T) {
		word, definition := "Test", "This is the original definition"
		dictionary := Dictionary{word: definition}

		err := dictionary.Delete(word)
		if err != nil {
			t.Errorf("Expected: nil - Actual: %s", err)
		}

		expected := ErrDictionaryWordNotFound(word)
		_, actual := dictionary.Search(word)
		if actual == nil {
			t.Errorf("Expected: %s - Actual: nil", expected)
		}
		if actual.Error() != expected.Error() {
			t.Errorf("Expected: %s - Actual: %s", expected, actual)
		}
	})
}
