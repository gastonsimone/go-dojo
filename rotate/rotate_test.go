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

func TestReverseEmpty(t *testing.T) {
	var a []string = nil

	Reverse(a)

	var want []string = nil
	if !areEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}
}

func TestReverseSingleItem(t *testing.T) {
	a := []string{"1"}

	Reverse(a)

	want := []string{"1"}
	if !areEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}
}

func TestReverseEven(t *testing.T) {
	a := []string{"1", "2", "3", "4"}

	Reverse(a)
	want := []string{"4", "3", "2", "1"}
	if !areEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}
}

func TestReverseOdd(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5"}

	Reverse(a)
	want := []string{"5", "4", "3", "2", "1"}
	if !areEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}
}

func TestRotateEmpty(t *testing.T) {
	err := Rotate(nil, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRotateOneLeft(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7"}

	err := Rotate(a, -1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []string{"2", "3", "4", "5", "6", "7", "1"}
	if !areEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}
}

func TestRotateNothing(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7"}

	err := Rotate(a, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []string{"1", "2", "3", "4", "5", "6", "7"}
	if !areEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}
}

func TestRotateOne(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7"}

	err := Rotate(a, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []string{"7", "1", "2", "3", "4", "5", "6"}
	if !areEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}
}

func TestRotateAllButOne(t *testing.T) {
	a := []string{"7", "1", "2", "3", "4", "5", "6"}

	err := Rotate(a, 6)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []string{"1", "2", "3", "4", "5", "6", "7"}
	if !areEqual(a, want) {
		t.Fatalf("got %v want %v", a, want)
	}
}

func TestRotateEverything(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7"}

	err := Rotate(a, 7)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []string{"1", "2", "3", "4", "5", "6", "7"}
	if !areEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}
}

func TestRotateMoreThanLength(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7"}

	err := Rotate(a, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []string{"5", "6", "7", "1", "2", "3", "4"}
	if !areEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}
}
