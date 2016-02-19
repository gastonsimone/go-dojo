// Package strarray implements simple functions to manipulate string arrays and
// slices
package strarray

// AreEqual returns true only if both string slices s and t are equal
func AreEqual(s, t []string) bool {
	if len(s) != len(t) {
		return false
	}

	for i, v := range s {
		if v != t[i] {
			return false
		}
	}

	return true
}

// reverses a slice in place.
func Reverse(a []string) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
