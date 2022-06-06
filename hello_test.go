package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, want, got interface{}) {
		t.Helper()
		if want != got {
			t.Errorf("want %q got %q", want, got)
		}
	}

	t.Run("greeting people", func(t *testing.T) {
		want := "Hello, Chris"
		got := Hello("Chris", "")
		assertCorrectMessage(t, want, got)
	})

	t.Run("greeting default", func(t *testing.T) {
		want := "Hello, World"
		got := Hello("", "")
		assertCorrectMessage(t, want, got)
	})

	t.Run("greeting in Spanish", func(t *testing.T) {
		want := "Hola, Elodie"
		got := Hello("Elodie", "Spanish")
		assertCorrectMessage(t, want, got)
	})
}
