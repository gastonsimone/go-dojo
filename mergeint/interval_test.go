package mergeint

import "testing"

func TestOverlaps1(t *testing.T) {
	i := Interval{1, 10}
	j := Interval{5, 15}
	if !i.Overlaps(j) {
		t.Fatalf("%v and %v do overlap", i, j)
	}
}

func TestOverlaps2(t *testing.T) {
	i := Interval{1, 10}
	j := Interval{10, 15}
	if !i.Overlaps(j) {
		t.Fatalf("%v and %v do overlap", i, j)
	}
}

func TestOverlaps3(t *testing.T) {
	i := Interval{10, 20}
	j := Interval{1, 10}
	if !i.Overlaps(j) {
		t.Fatalf("%v and %v do overlap", i, j)
	}
}

func TestOverlaps4(t *testing.T) {
	i := Interval{1, 10}
	j := Interval{11, 20}
	if i.Overlaps(j) {
		t.Fatalf("%v and %v do not overlap", i, j)
	}
}

func TestMerge1(t *testing.T) {
	i := Interval{1, 10}
	j := Interval{5, 15}
	i.Merge(j)
	want := Interval{1, 15}
	if i != want {
		t.Fatalf("got %v, want %v", i, want)
	}
}

func TestMerge2(t *testing.T) {
	i := Interval{5, 15}
	j := Interval{1, 10}
	i.Merge(j)
	want := Interval{1, 15}
	if i != want {
		t.Fatalf("got %v, want %v", i, want)
	}
}

func TestMergeNoEverlap(t *testing.T) {
	i := Interval{1, 10}
	j := Interval{11, 20}
	i.Merge(j)
	want := Interval{1, 10}
	if i != want {
		t.Fatal("unexpected change merging non-overlapping intervals")
	}
}
