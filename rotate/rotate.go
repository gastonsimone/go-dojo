// Provides function to rotate slices in place, using O(1) and
// O(n) time, with n = the slice length.
package rotate

import "github.com/gastonsimone/go-dojo/strarray"

// rotates slice in place k positions
// Memory: O(1)
// Time: O(n), n = len(array)
func Rotate(s []string, k int) {
	if len(s) <= 0 {
		// Nothing to do
		return
	}

	// Do not rotate more than necessary
	k = k % len(s)

	switch {
	case k == 0:
		// Nothing to do
		return
	case k < 0:
		// Convert a left-rotate to a right-rotate
		k += len(s)
	}

	split := len(s) - k
	strarray.Reverse(s[:split])
	strarray.Reverse(s[split:])
	strarray.Reverse(s)
	return
}
