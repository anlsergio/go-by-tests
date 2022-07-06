package store

import (
	"github.com/anlsergio/go-by-tests/webapp/tests"
	"io"
	"testing"
)

func TestTape_Write(t *testing.T) {
	file, cleanDB := tests.CreateTempFile(t, "1234")
	defer cleanDB()

	tape := &tape{file}

	tape.Write([]byte("abc"))

	file.Seek(0, 0)
	newFileContents, _ := io.ReadAll(file)

	want := "abc"
	got := string(newFileContents)

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
