package maps

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("could not find the word you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	// Check if word already exist
	_, err := d.Search(word)

	// Handle response error
	switch err {
	case ErrNotFound:
		// Normal case: word does not exist so insert 'new' definition
		d[word] = definition
	case nil:
		// nil Error means: requested word already exist in Dictionary
		// We can not overwrite existing word-definition
		return ErrWordExists
	}

	// word added successfully (no error)
	return nil
}
