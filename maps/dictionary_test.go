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
