// Package median provides a solution for the "Efficient median of two sorted
// arrays" problem.
package median

import (
	"errors"
)

// Median calculates the median of the two sorted arrays s and t. The overall
// run time complexity is O(log(min(m,n))) on average and O(log(m)+log(n))
// worst case, where m = len(s) and n = len(t).
func Median(s, t []string) (string, error) {
	if len(s) == 0 {
		if len(t) == 0 {
			return "", errors.New("arrays cannot be both empty")
		}

		// s is empty, general median is the median of t
		return t[len(t)/2], nil
	}

	if len(t) == 0 {
		// t is empty, general median is the median of s
		return s[len(s)/2], nil
	}

	// Scan the smaller slice first.
	var small, big []string

	if len(s) <= len(t) {
		small = s
		big = t
	} else {
		small = t
		big = s
	}

	if m, found := medianScan(small, big); found {
		return m, nil
	}

	m, _ := medianScan(big, small)
	return m, nil
}

type result uint

const (
	greater = iota
	lower
)

// medianScan scans only slice s for the overall median. Therefore it is better
// to call this function with the smallest slice first, which is what
// medianDifferentSize does. Time complexity: O(log(len(s))) worst case.
func medianScan(s, t []string) (median string, found bool) {
	// total is the total number of elements of s and t combined
	total := len(s) + len(t)

	// high is the number of elements that have been confirmed so far to be
	// greater than the overall median.
	high := 0

	// Optimization: if the overall median is in s, it will not be further from
	// the middle of s than the number of elements in t. So we only need to
	// analyse the middle section of s.
	if len(s) > len(t) && len(t) > 2 {
		from := len(s)/2 - len(t)/2
		to := len(s)/2 + len(t)/2
		high += len(s) - to - 1
		s = s[from : to+1]
	}

	for len(s) > 1 {
		// index of the current guess in s
		is := len(s) / 2

		// guess is the element check to be the overall median. On the first
		// iteration, guess is the median of array s.
		guess := s[is]

		// if guess is the overall median, scover is the number of elements in
		// s that are greater than the overall median.
		scover := len(s) - is - 1

		// if guess is the overall median, tcover is the number of elements in
		// t that are greater than the overall median.
		tcover := (total / 2) - (scover + high)

		// check is the position in t that needs to be checked to confirm guess
		// is correct
		check := len(t) - tcover

		var r result
		switch {
		case check < 0:
			r = greater
		case check >= len(t):
			r = lower
		case guess <= t[check]:
			if check == 0 || t[check-1] <= guess {
				return guess, true
			}
			r = lower
		default:
			r = greater
		}

		if r == greater {
			// guess is greater than the overall median
			// so the lower half of s must be searched
			s = s[:is]
			high += (scover + 1)
		} else {
			// guess is lower than the overall median
			// so the upper half of s must be searched
			s = s[is:]
		}
	}

	// border case: len(s) == 1
	guess := s[0]
	tcover := (total / 2) - high
	check := len(t) - tcover

	if (check < len(t) && guess <= t[check]) || tcover == 0 {
		if check == 0 || t[check-1] <= guess {
			return guess, true
		}
	}

	return "", false
}
