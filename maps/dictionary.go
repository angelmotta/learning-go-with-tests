package maps

var (
	// We can use DictionaryErr as a type Error (because this type implements Error interface)
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
	// Instead of explicitly import errors package and create a New Error
	//ErrWordExists = errors.New("cannot add word because it already exists")
)

// DictionaryErr type are errors that can happen when interacting with the dictionary
// DictionaryErr is type string which also implements the Error Interface
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary type is our custom hashmap used a dictionary
type Dictionary map[string]string

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
		// Happy path: word does not exist, so we can Add 'new' definition
		d[word] = definition
	case nil:
		// nil Error means: requested word already exist in Dictionary
		// We can not overwrite existing word-definition
		return ErrWordExists
	}

	// word added successfully (no error)
	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		// Tried to update a word not defined in the Dictionary
		return ErrWordDoesNotExist
	case nil:
		// Happy path: Word-definition exist in the dictionary, so we can update this
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
