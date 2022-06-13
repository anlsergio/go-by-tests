package reflection__test

import (
	reflection_ "hello/reflection"
	"testing"
)

func TestWalk(t *testing.T) {
	want := "Chris"
	x := struct {
		Name string
	}{want}

	var calledStrings []string

	spyFunc := func(input string) {
		calledStrings = append(calledStrings, input)
	}

	reflection_.Walk(x, spyFunc)

	if len(calledStrings) != 1 {
		t.Errorf("wrong number of function calls, want %d got %d", 1, len(calledStrings))
	}

	got := calledStrings[0]

	if want != got {
		t.Errorf("want %q got %q", want, got)
	}
}
