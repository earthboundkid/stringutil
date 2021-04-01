package stringutil

// CharMap is an array for use in mapping from byte to bool
// and testing whether the characters of string are in range.
type CharMap = [256]bool

func MakeCharMap(s string) *CharMap {
	var m CharMap
	for _, c := range []byte(s) {
		m[c] = true
	}
	return &m
}

func InRange(s string, m *CharMap) bool {
	for _, b := range []byte(s) {
		if !m[b] {
			return false
		}
	}
	return true
}

// CharBitField is a bit field for testing character ranges.
// It is not as fast as CharMap, but it is 8 times smaller
// if that's important for some reason.
type CharBitField = [256 / 8]byte

func MakeCharBitField(s string) *CharBitField {
	var m CharBitField
	for _, c := range []byte(s) {
		i := c / 8
		mask := byte(1 << (c % 8))
		m[i] |= mask
	}
	return &m
}

func InBitField(s string, m *CharBitField) bool {
	for _, c := range []byte(s) {
		i := c / 8
		mask := byte(1 << (c % 8))
		if m[i]&mask == 0 {
			return false
		}
	}
	return true
}
