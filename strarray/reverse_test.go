package strarray

import (
	"testing"
)

func TestReverseEmpty(t *testing.T) {
	var a []string = nil

	Reverse(a)

	var want []string = nil
	if !AreEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}
}

func TestReverseSingleItem(t *testing.T) {
	a := []string{"1"}

	Reverse(a)

	want := []string{"1"}
	if !AreEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}
}

func TestReverseEven(t *testing.T) {
	a := []string{"1", "2", "3", "4"}

	Reverse(a)
	want := []string{"4", "3", "2", "1"}
	if !AreEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}
}

func TestReverseOdd(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5"}

	Reverse(a)
	want := []string{"5", "4", "3", "2", "1"}
	if !AreEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}
}
