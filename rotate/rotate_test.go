package rotate

import (
	"github.com/gastonsimone/go-dojo/strarray"
	"testing"
)

func TestRotateEmpty(t *testing.T) {
	Rotate(nil, 1)
}

func TestRotateOneLeft(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7"}

	Rotate(a, -1)

	want := []string{"2", "3", "4", "5", "6", "7", "1"}
	if !strarray.AreEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}
}

func TestRotateNothing(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7"}

	Rotate(a, 0)

	want := []string{"1", "2", "3", "4", "5", "6", "7"}
	if !strarray.AreEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}
}

func TestRotateOne(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7"}

	Rotate(a, 1)

	want := []string{"7", "1", "2", "3", "4", "5", "6"}
	if !strarray.AreEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}
}

func TestRotateAllButOne(t *testing.T) {
	a := []string{"7", "1", "2", "3", "4", "5", "6"}

	Rotate(a, 6)

	want := []string{"1", "2", "3", "4", "5", "6", "7"}
	if !strarray.AreEqual(a, want) {
		t.Fatalf("got %v want %v", a, want)
	}
}

func TestRotateEverything(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7"}

	Rotate(a, 7)

	want := []string{"1", "2", "3", "4", "5", "6", "7"}
	if !strarray.AreEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}
}

func TestRotateMoreThanLength(t *testing.T) {
	a := []string{"1", "2", "3", "4", "5", "6", "7"}

	Rotate(a, 10)

	want := []string{"5", "6", "7", "1", "2", "3", "4"}
	if !strarray.AreEqual(a, want) {
		t.Fatalf("got %v, want %v", a, want)
	}
}
