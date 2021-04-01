package stringutil_test

import (
	"testing"

	"github.com/carlmjohnson/stringutil"
)

func assertBool(t *testing.T, want, got bool) {
	t.Helper()
	assert(t, want == got, "got %t", got)
}
func TestCharMap(t *testing.T) {
	cases := map[string]struct {
		chars, s string
		ok       bool
	}{
		"empty":      {"", "", true},
		"one":        {"a", "aaaa", true},
		"blank":      {"", "aaaa", false},
		"missing":    {"a", "aba", false},
		"multi":      {"abcd", "abababa", true},
		"badspread":  {"\x00\xff", "abc", false},
		"goodspread": {"\x001a\xff", "\x001a\xff\x001a\xff", true},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			m := stringutil.MakeCharMap(tc.chars)
			ok := stringutil.InRange(tc.s, m)
			assertBool(t, tc.ok, ok)
		})
	}
}

func TestCharBitField(t *testing.T) {
	cases := map[string]struct {
		chars, s string
		ok       bool
	}{
		"empty":      {"", "", true},
		"one":        {"a", "aaaa", true},
		"blank":      {"", "aaaa", false},
		"missing":    {"a", "aba", false},
		"multi":      {"abcd", "abababa", true},
		"badspread":  {"\x00\xff", "abc", false},
		"goodspread": {"\x001a\xff", "\x001a\xff\x001a\xff", true},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			m := stringutil.MakeCharBitField(tc.chars)
			ok := stringutil.InBitField(tc.s, m)
			assertBool(t, tc.ok, ok)
		})
	}
}

var bsink bool

func BenchmarkCharMap(b *testing.B) {
	ss := []string{
		"the quick brown fox jumps over a lazy dog 12345",
		"the quick brown fox \njumps over a lazy dog 12345",
		"the quick brown fox jumps over a lazy dog???",
		"the quick brown fox jumps over a lazy dog!",
	}
	m := stringutil.MakeCharMap("abcdefghijlkmnopqrstuvwxyz1234567890 ")
	for i := 0; i < b.N; i++ {
		s := ss[i%len(ss)]
		bsink = stringutil.InRange(s, m)
	}
}

func BenchmarkCharBitField(b *testing.B) {
	ss := []string{
		"the quick brown fox jumps over a lazy dog 12345",
		"the quick brown fox \njumps over a lazy dog 12345",
		"the quick brown fox jumps over a lazy dog???",
		"the quick brown fox jumps over a lazy dog!",
	}
	m := stringutil.MakeCharBitField("abcdefghijlkmnopqrstuvwxyz1234567890 ")
	for i := 0; i < b.N; i++ {
		s := ss[i%len(ss)]
		bsink = stringutil.InBitField(s, m)
	}
}
