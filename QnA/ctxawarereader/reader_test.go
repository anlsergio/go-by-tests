package ctxawarereader

import (
	"context"
	"strings"
	"testing"
)

func TestContextAwareReader(t *testing.T) {
	t.Run("lets just see how a normal reader works", func(t *testing.T) {
		rdr := strings.NewReader("ABCDEF")

		got := make([]byte, 3)

		_, err := rdr.Read(got)
		if err != nil {
			t.Fatal(err)
		}

		assertBufferHas(t, "ABC", got)

		_, err = rdr.Read(got)
		if err != nil {
			t.Fatal(err)
		}

		assertBufferHas(t, "DEF", got)
	})

	t.Run("behaves like a normal reader", func(t *testing.T) {
		rdr := NewCancellableReader(context.Background(), strings.NewReader("ABCDEF"))

		got := make([]byte, 3)

		_, err := rdr.Read(got)
		if err != nil {
			t.Fatal(err)
		}

		assertBufferHas(t, "ABC", got)

		_, err = rdr.Read(got)
		if err != nil {
			t.Fatal(err)
		}

		assertBufferHas(t, "DEF", got)
	})

	t.Run("stops reading when cancelled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())

		rdr := NewCancellableReader(ctx, strings.NewReader("ABCDEF"))

		got := make([]byte, 3)

		_, err := rdr.Read(got)
		if err != nil {
			t.Fatal(err)
		}

		assertBufferHas(t, "ABC", got)

		cancel()

		n, err := rdr.Read(got)
		if err == nil {
			t.Error("expected an error after cancellation but didn't get any")
		}

		if n > 0 {
			t.Errorf("expected 0 bytes to be read after cancellation, but got %d", n)
		}
	})
}

func assertBufferHas(t testing.TB, want string, buf []byte) {
	t.Helper()

	got := string(buf)
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
