package maps_test

import (
	"hello/maps"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := maps.Dictionary{
		"test": "nothing to see here",
	}

	t.Run("known word", func(t *testing.T) {
		want := "nothing to see here"
		got, _ := dictionary.Search("test")

		assertStrings(t, want, got)
	})

	t.Run("unknown word", func(t *testing.T) {
		want := maps.ErrNotFound
		_, err := dictionary.Search("unknown")
		if err == nil {
			t.Fatal("error expected but didn't get any")
		}

		assertError(t, want, err)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := maps.Dictionary{}
		word := "test"
		definition := "nothing to see here"
		err := dictionary.Add(word, definition)

		assertError(t, nil, err)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "nothing to see here"
		dictionary := maps.Dictionary{word: definition}
		err := dictionary.Add(word, "[overwritten]")

		assertError(t, maps.ErrWordExists, err)
		assertDefinition(t, dictionary, word, definition)
	})
}

func assertStrings(t testing.TB, want string, got string) {
	t.Helper()

	if want != got {
		t.Errorf("want %q got %q given: %q", want, got, "test")
	}
}

func assertError(t testing.TB, want, got error) {
	t.Helper()

	if want != got {
		t.Errorf("want error %q got %q", want, got)
	}
}

func assertDefinition(t testing.TB, dictionary maps.Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("unexpected error: ", err)
	}

	if definition != got {
		t.Errorf("want %q got %q", definition, got)
	}
}
