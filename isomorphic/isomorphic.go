package isomorphic

import "unicode/utf8"

// AreIsomorphic returns true only if strings s and t are isomorphic.
// Two strings are isomorphic if you can get one of them by replacing all the
// characters in the other.
func AreIsomorphic(s, t string) bool {
	slen := utf8.RuneCountInString(s)
	tlen := utf8.RuneCountInString(t)

	// Strings are not isomorphic if they have different lengths.
	if slen != tlen {
		return false
	}

	// Simple border case for efficiency. At this point we know both strings
	// have the same length, so it is enough to check just one of them.
	if slen == 0 {
		return true
	}

	charmap := make(map[rune]rune)

	for slen > 0 {
		rs, ssize := utf8.DecodeRuneInString(s)
		rt, tsize := utf8.DecodeRuneInString(t)

		r, found := charmap[rs]
		if found {
			if r != rt {
				return false
			}
		} else {
			charmap[rs] = rt
		}

		s = s[ssize:]
		slen -= ssize

		t = t[tsize:]
		//tlen -= tsize // Not really necessary, just for completeness
	}

	return true
}
