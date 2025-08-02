package maps

import (
	"fmt"
	"testing"
)

var (
	word       = "test"
	definition = "this is test"
	dictionary = Dictionary{word: definition}
)

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is test"

		assertStrings(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, got := dictionary.Search("test1")
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		word := "new test"
		definition := "this is new test"

		err := dictionary.Add(word, definition)
		fmt.Println(dictionary)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {

		word := "test"
		definition := "this is test"

		err := dictionary.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("update existing word", func(t *testing.T) {
		word := "test"
		newDefinition := "this"
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("update not existing word", func(t *testing.T) {
		word := "test2"
		newDefinition := "this"
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, ErrorWordNotExists)
		assertDefinition(t, dictionary, word, newDefinition)
		fmt.Println(dictionary)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary.Delete(word)
	_, err := dictionary.Search(word)
	assertError(t, err, ErrNotFound)
}

// this is test helper function
func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {

	t.Helper()
	value, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if value != definition {
		t.Errorf("got '%s' but want '%s'", value, definition)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got is '%s' but want '%s'", got, want)
	}
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}
