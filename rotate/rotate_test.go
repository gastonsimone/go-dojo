package rotate

import (
	"testing"
)

func areEqual(s, t []string) bool {
	if len(s) != len(t) {
		return false
	}

	for i, v := range s {
		if v != t[i] {
			return false
		}
	}

	return true
}

func TestReverse(t *testing.T) {
	var a []string = nil

	Reverse(a)
	var want []string = nil
	if !areEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}

	a = []string{"1"}

	Reverse(a)
	want = []string{"1"}
	if !areEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}

	a = []string{"1", "2", "3", "4"}

	Reverse(a)
	want = []string{"4", "3", "2", "1"}
	if !areEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}

	a = []string{"1", "2", "3", "4", "5"}

	Reverse(a)
	want = []string{"5", "4", "3", "2", "1"}
	if !areEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}
}

func TestRotate(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7"}

	Rotate(a, 0)
	want1 := []string{"1", "2", "3", "4", "5", "6", "7"}
	if !areEqual(a, want1) {
		t.Fatalf("got %v, want %v", a, want1)
	}

	Rotate(a, 1)
	want2 := []string{"7", "1", "2", "3", "4", "5", "6"}
	if !areEqual(a, want2) {
		t.Fatalf("got %v, want %v", a, want2)
	}

	Rotate(a, 6)
	if !areEqual(a, want1) {
		t.Fatalf("got %v, want %v", a, want1)
	}

	Rotate(a, 7)
	if !areEqual(a, want1) {
		t.Fatalf("got %v, want %v", a, want1)
	}

	Rotate(a, 10)
	want3 := []string{"5", "6", "7", "1", "2", "3", "4"}
	if !areEqual(a, want3) {
		t.Fatalf("got %v, want %v", a, want3)
	}
}
