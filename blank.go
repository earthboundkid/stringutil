package stringutil

// First returns the first non-blank string.
func First(ss ...string) string {
	for _, s := range ss {
		if s != "" {
			return s
		}
	}
	return ""
}

// Last returns the last non-blank string.
func Last(ss ...string) string {
	for i := range ss {
		ri := len(ss) - 1 - i
		s := ss[ri]
		if s != "" {
			return s
		}
	}
	return ""
}

// HasString returns true if an equal string is in the slice.
func HasString(s string, ss ...string) bool {
	for _, s2 := range ss {
		if s == s2 {
			return true
		}
	}
	return false
}

// HasBlank returns true if any string is blank.
func HasBlank(ss ...string) bool {
	return HasString("", ss...)
}
