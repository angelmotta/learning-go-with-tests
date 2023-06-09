package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		// In this test we are interested verifying the 'value'
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		// In this test we are interested verifying the 'error'
		_, gotErr := dictionary.Search("unknown")

		// Check if err is null before pass it to Assertion and call err.Error()
		if gotErr == nil {
			t.Fatal("expected to get an error.")
		}

		// err is not Null, so we can call err.Error()
		assertError(t, gotErr, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test definition"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test definition")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test definition"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		// Invoke Update method
		receivedErr := dictionary.Update(word, newDefinition)

		// Verify Update method
		assertError(t, receivedErr, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		updateDefinition := "this is just a test"
		dictionary := Dictionary{}

		// Invoke Update method
		receivedErr := dictionary.Update(word, updateDefinition)

		// Verify Update method
		assertError(t, receivedErr, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

// assertDefinition verify a definition exist for a given word and the given definition (got) is the expected (parameter)
func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	gotDefinition, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	assertStrings(t, gotDefinition, definition)
}

// assertStrings check both strings "got" and "want" are the same
func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
