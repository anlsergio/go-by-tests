package tests

import (
	"os"
	"testing"
)

func CreateTempFile(t testing.TB, initialData string) (fileBuffer *os.File, removeTempFile func()) {
	t.Helper()

	tempFile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}

	tempFile.Write([]byte(initialData))

	removeTempFile = func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}

	return tempFile, removeTempFile
}

func AssertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}
