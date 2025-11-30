package maps

import (
	"errors"
	"fmt"
)

var (
	ErrDictionaryWordNotFound = func(word string) error {
		return errors.New(fmt.Sprintf("word %s not found in the dictionary", word))
	}
	ErrDictionaryWordAlreadyExists = func(word string) error {
		return errors.New(fmt.Sprintf("word %s already exists in the dictionary", word))
	}
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrDictionaryWordNotFound(word)
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	if err != nil {
		d[word] = definition
		return nil
	}

	return ErrDictionaryWordAlreadyExists(word)
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	if err != nil {
		return err
	}

	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err != nil {
		return err
	}

	delete(d, word)
	return nil
}
