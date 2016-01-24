// Provides functions to rotate and reverse slices in place, using O(1) and
// O(n) time, with n = the slice length.
package rotate

import (
	"errors"
)

// rotates slice in place k positions
// Memory: O(1)
// Time: O(n), n = len(array)
func Rotate(s []string, k int) error {
	if k < 0 {
		return errors.New("invalid number of positions to rotate")
	}

	if len(s) <= 0 {
		return errors.New("empty s")
	}

	// Do not rotate more than necessary
	k = k % len(s)

	split := len(s) - k
	Reverse(s[:split])
	Reverse(s[split:])
	Reverse(s)
	return nil
}

// reverses a slice in place.
func Reverse(a []string) {
	i, j := 0, len(a)-1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}
