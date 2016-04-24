package mergeint

import "sort"

// IntervalSet represents a set of Intervals.
type IntervalSet []Interval

// Contains looks for an interval in a sorted list of intervals
func (s *IntervalSet) Contains(i Interval) bool {
	for _, j := range *s {
		if j == i {
			return true
		}
	}
	return false
}

// Using defines a function that merges the intervals in a set
type Using func(s *IntervalSet) IntervalSet

// Merge merges all the intervals in s using the corresponding implementation
func (merge Using) Merge(s *IntervalSet) IntervalSet {
	return merge(s)
}

// ByStart implements sort.Interface for IntervalSet on the start field
type ByStart IntervalSet

func (s ByStart) Len() int           { return len(s) }
func (s ByStart) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByStart) Less(i, j int) bool { return s[i].start < s[j].start }

// SortFirst builds a new IntervalSet, result of merging all the overlapping
// intervals in s. This solution sorts s before starting, and then merges all
// the intervals. The initial sort ensures only the last merged interval needs
// to be checked for overlapping, which reduces the number of comparisons that
// need to be performed.
func SortFirst(s *IntervalSet) IntervalSet {
	sort.Sort(ByStart(*s))
	merged := make(IntervalSet, 0, len(*s))
	merged = append(merged, (*s)[0])
	for i := 1; i < len(*s); i++ {
		last := len(merged) - 1
		if merged[last].Overlaps((*s)[i]) {
			merged[last].Merge((*s)[i])
		} else {
			merged = append(merged, (*s)[i])
		}
	}
	return merged
}

// NoSort builds a new IntervalSet, result of merging all the overlapping
// intervals in s. This solution does not need to sort the intervals before
// merging, but it needs to check all the already-merged intervals with every
// new interval it has to process (this is why it has a nested for loop). It
// also removes and re-inserts elements in the merged set as it finds more
// intervals that can be merged.
func NoSort(s *IntervalSet) IntervalSet {
	merged := make(IntervalSet, 0, len(*s))
	for _, ints := range *s {
		for i := 0; i < len(merged); {
			intm := merged[i]
			if ints.Overlaps(intm) {
				ints.Merge(intm)
				// remove position i (intm) from merged set
				merged = append(merged[:i], merged[i+1:]...)
			} else {
				i++
			}
		}
		merged = append(merged, ints)
	}
	return merged
}
