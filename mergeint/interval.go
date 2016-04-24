package mergeint

// Interval represents an interval in the set of integers.
// start should be smaller or equal to end.
type Interval struct {
	start int
	end   int
}

// Overlaps returns true only if intervals i and j overlap
func (i *Interval) Overlaps(j Interval) bool {
	return (i.start <= j.end && j.start <= i.end)
}

// Merge merges interval j on i, overwriting i
func (i *Interval) Merge(j Interval) {
	if !i.Overlaps(j) {
		return
	}

	if j.start < i.start {
		i.start = j.start
	}

	if j.end > i.end {
		i.end = j.end
	}
}
