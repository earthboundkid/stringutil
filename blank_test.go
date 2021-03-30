package stringutil_test

import (
	"strings"
	"testing"

	"github.com/carlmjohnson/stringutil"
)

func TestFirst(t *testing.T) {
	cases := map[string]struct {
		input []string
		want  string
	}{
		"empty":  {[]string{}, ""},
		"one":    {strings.Fields("a"), "a"},
		"two":    {strings.Fields("b a"), "b"},
		"blank":  {strings.Split(",a", ","), "a"},
		"blanks": {strings.Split(",,,a,,b", ","), "a"},
	}
	for name, tc := range cases {
		t.Run(name+"-nil", func(t *testing.T) {
			got := stringutil.First(tc.input...)
			assert(t, got == tc.want, "want %v; got %v", tc.want, got)
		})
	}
}
