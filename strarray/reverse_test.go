package strarray

import (
	"testing"
)

func TestReverseEmpty(t *testing.T) {
	var a []string

	Reverse(a)

	var want []string
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

func BenchmarkReverse(b *testing.B) {
	a := []string{"1", "2", "3", "4", "5", "7", "8", "9", "10", "11", "12"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Reverse(a)
	}
}
