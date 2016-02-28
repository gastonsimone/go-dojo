package strarray

import (
	"testing"
)

func TestAreEqualNil(t *testing.T) {
	if !AreEqual(nil, nil) {
		t.Fatal("nil slices not considered equal")
	}
}

func TestAreEqualNilNotNil(t *testing.T) {
	a := []string{"one"}
	if AreEqual(a, nil) {
		t.Fatalf("%v equal to nil", a)
	}

	if AreEqual(nil, a) {
		t.Fatalf("nil equal to %v", a)
	}
}

func TestAreEqualDifferentLength(t *testing.T) {
	a := []string{"one"}
	b := []string{"one", "two"}
	if AreEqual(a, b) {
		t.Fatalf("%v equal to %v", a, b)
	}
}

func TestAreEqualDifferentCase(t *testing.T) {
	a := []string{"one", "two"}
	b := []string{"one", "Two"}
	if AreEqual(a, b) {
		t.Fatalf("%v equal to %v", a, b)
	}
}

func TestAreEqual(t *testing.T) {
	a := []string{"one", "two", "three", "four"}
	b := []string{"one", "two", "three", "four"}
	if !AreEqual(a, b) {
		t.Fatalf("%v not equal to %v", a, b)
	}
}

func BenchmarkAreEqual(b *testing.B) {
	s := []string{"one", "two", "three", "four", "five", "six", "seven", "eight"}
	t := []string{"one", "two", "three", "four", "five", "six", "seven", "eight"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AreEqual(s, t)
	}
}
