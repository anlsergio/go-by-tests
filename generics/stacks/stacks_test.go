package stacks_test

import (
	"hello/generics/stacks"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(stacks.Stack[int])

		// check stack is empty
		AssertTrue(t, myStackOfInts.IsEmpty())

		// add a thing, then check it's not empty
		const firstValue = 123
		myStackOfInts.Push(firstValue)
		AssertFalse(t, myStackOfInts.IsEmpty())

		// add another thing, pop it back again
		const secondValue = 456
		myStackOfInts.Push(secondValue)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, secondValue, value)
		value, _ = myStackOfInts.Pop()
		AssertEqual(t, firstValue, value)
		AssertTrue(t, myStackOfInts.IsEmpty())

		// can get the numbers we put in as numbers, not untyped interface{}
		myStackOfInts.Push(1)
		myStackOfInts.Push(2)
		firstNum, _ := myStackOfInts.Pop()
		secondNum, _ := myStackOfInts.Pop()
		AssertEqual(t, firstNum+secondNum, 3)
	})

	t.Run("string stack", func(t *testing.T) {
		myStackOfStrings := new(stacks.Stack[string])

		// check stack is empty
		AssertTrue(t, myStackOfStrings.IsEmpty())

		// add a thing, then check it's not empty
		myStackOfStrings.Push("123")
		AssertFalse(t, myStackOfStrings.IsEmpty())

		// add another thing, pop it back again
		myStackOfStrings.Push("456")
		value, _ := myStackOfStrings.Pop()
		AssertEqual(t, value, "456")
		value, _ = myStackOfStrings.Pop()
		AssertEqual(t, value, "123")
		AssertTrue(t, myStackOfStrings.IsEmpty())
	})

	//t.Run("interface stack dx is horrid as opposed to generics", func(t *testing.T) {
	//	myStackOfInts := new(stacks.StackOfInts)
	//
	//	myStackOfInts.Push(1)
	//	myStackOfInts.Push(2)
	//	firstNum, _ := myStackOfInts.Pop()
	//	secondNum, _ := myStackOfInts.Pop()
	//
	//	// get our ints from out interface{}
	//	reallyFirstNum, ok := firstNum.(int)
	//	AssertTrue(t, ok) // need to check we definitely got an int out of the interface{}
	//
	//	reallySecondNum, ok := secondNum.(int)
	//	AssertTrue(t, ok) // and again!
	//
	//	AssertEqual(t, reallyFirstNum+reallySecondNum, 3)
	//})
}

func AssertEqual[T comparable](t *testing.T, want, got T) {
	t.Helper()
	if want != got {
		t.Errorf("want %v, got %v", want, got)
	}
}

func AssertNotEqual[T comparable](t *testing.T, a, b T) {
	t.Helper()
	if a == b {
		t.Errorf("expected %+v not to be equal to %+v", a, b)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Error("want true, got ", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Error("want false, got ", got)
	}
}
