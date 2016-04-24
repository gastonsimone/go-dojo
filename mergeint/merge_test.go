package mergeint

import "testing"

func validateSet(strategy string, t *testing.T, merged, want IntervalSet) {
	if len(merged) != len(want) {
		t.Fatalf("%v: unexpected number of merged intervals:\nGot  %v: %v\nWant %v: %v\n",
			strategy, len(merged), merged, len(want), want)
	}

	for _, interval := range want {
		if !merged.Contains(interval) {
			t.Fatalf("%v: interval %v not found in %v\n", strategy, interval, merged)
		}
	}
}

func TestContainsFound(t *testing.T) {
	s := IntervalSet{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	i := Interval{8, 10}
	if !s.Contains(i) {
		t.Fatal("interval not found")
	}
}

func TestContainsNotFound(t *testing.T) {
	s := IntervalSet{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	i := Interval{20, 30}
	if s.Contains(i) {
		t.Fatal("interval found")
	}
}

func TestMerge1IntervalsBasic(t *testing.T) {
	s := IntervalSet{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	want := IntervalSet{{1, 6}, {8, 10}, {15, 18}}

	merged := Using(NoSort).Merge(&s)
	validateSet("NoSort", t, merged, want)

	merged = Using(SortFirst).Merge(&s)
	validateSet("SortFirst", t, merged, want)
}

func TestMerge1IntervalsBasicUnsorted(t *testing.T) {
	s := IntervalSet{{2, 6}, {15, 18}, {8, 10}, {1, 3}}
	want := IntervalSet{{1, 6}, {8, 10}, {15, 18}}

	merged := Using(NoSort).Merge(&s)
	validateSet("NoSort", t, merged, want)

	merged = Using(SortFirst).Merge(&s)
	validateSet("SortFirst", t, merged, want)
}

func TestMergeAll(t *testing.T) {
	s := IntervalSet{{1, 3}, {2, 6}, {5, 10}, {10, 18}}
	want := IntervalSet{{1, 18}}

	merged := Using(NoSort).Merge(&s)
	validateSet("NoSort", t, merged, want)

	merged = Using(SortFirst).Merge(&s)
	validateSet("SortFirst", t, merged, want)
}

func TestMergeAllUnsorted(t *testing.T) {
	s := IntervalSet{{2, 6}, {1, 3}, {10, 18}, {5, 10}}
	want := IntervalSet{{1, 18}}

	merged := Using(NoSort).Merge(&s)
	validateSet("NoSort", t, merged, want)

	merged = Using(SortFirst).Merge(&s)
	validateSet("SortFirst", t, merged, want)
}

func TestMergeNothing(t *testing.T) {
	s := IntervalSet{{1, 3}, {4, 6}, {7, 10}, {11, 18}}

	merged := Using(NoSort).Merge(&s)
	validateSet("NoSort", t, merged, s)

	merged = Using(SortFirst).Merge(&s)
	validateSet("SortFirst", t, merged, s)
}

func BenchmarkNoSortMergeBasic(b *testing.B) {
	s := IntervalSet{{2, 6}, {15, 18}, {8, 10}, {1, 3}}
	merger := Using(NoSort)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		merger.Merge(&s)
	}
}

func BenchmarkSortFirstMergeBasic(b *testing.B) {
	s := IntervalSet{{2, 6}, {15, 18}, {8, 10}, {1, 3}}
	merger := Using(SortFirst)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		merger.Merge(&s)
	}
}

func BenchmarkNoSortMergeAll(b *testing.B) {
	s := IntervalSet{{2, 6}, {1, 3}, {10, 18}, {5, 10}}
	merger := Using(NoSort)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		merger.Merge(&s)
	}
}

func BenchmarkSortFirstMergeAll(b *testing.B) {
	s := IntervalSet{{2, 6}, {1, 3}, {10, 18}, {5, 10}}
	merger := Using(SortFirst)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		merger.Merge(&s)
	}
}
