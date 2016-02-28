package isomorphic

import "testing"

func TestEmpty(t *testing.T) {
	if !AreIsomorphic("", "") {
		t.Fatal("nil not considered isomorphic with itself")
	}
}

func TestSameString(t *testing.T) {
	if !AreIsomorphic("aaa", "aaa") {
		t.Fatal("'aaa' not considered isomorphic with itself")
	}
}

func TestDifferentLength(t *testing.T) {
	if AreIsomorphic("aaa", "bbbb") {
		t.Fatal("'aaa' considered isomorphic with 'bbbb'")
	}
}

func TestIsomorphic(t *testing.T) {
	if !AreIsomorphic("egg", "add") {
		t.Fatal("'egg' not considered isomorphic with 'add'")
	}
}

func TestNonIsomorphic(t *testing.T) {
	if AreIsomorphic("foo", "bar") {
		t.Fatal("'foo' considered isomorphic with 'bar'")
	}
}

func BenchmarkIsomorphic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AreIsomorphic("isomorphic", "economizer")
	}
}
